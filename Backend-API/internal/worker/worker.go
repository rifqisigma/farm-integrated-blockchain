package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Worker struct {
	farm         repository.FarmerRepository
	distribution repository.DistributorRepository
	seller       repository.SellerRepository
	collector    repository.CollectorRepository
	processor    repository.ProcessorRepository
	Redis        *redis.Client
	ticker       *time.Ticker
	quit         chan struct{}
	running      bool
	mu           sync.Mutex
}

func NewWorker(farm repository.FarmerRepository, distribution repository.DistributorRepository, seller repository.SellerRepository, collector repository.CollectorRepository, processor repository.ProcessorRepository, redis *redis.Client) *Worker {
	return &Worker{
		farm:         farm,
		distribution: distribution,
		seller:       seller,
		Redis:        redis,
		collector:    collector,
		processor:    processor,
	}
}

func (w *Worker) flushPendingItems() {
	keys, err := w.Redis.Keys(ctx, "behind:pending:*").Result()

	if err != nil {
		log.Fatalf("err redis queue: %s", err)
	}

	for _, key := range keys {
		fields, _ := w.Redis.HGetAll(ctx, key).Result()
		if len(fields) == 0 {
			continue
		}

		for _, value := range fields {
			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(value), &parsed); err != nil {
				log.Println("Failed to parse value:", err)
				continue
			}

			op, ok := parsed["op"].(string)
			if !ok {
				continue
			}

			dataMap, ok := parsed["data"].(map[string]interface{})
			if !ok {
				log.Println("Invalid data field")
				continue
			}

			switch op {
			case "bc_harvest":
				jsonValue, _ := json.Marshal(dataMap)
				errChan := make(chan error, 1)
				go func() {
					errChan <- w.HitBesuAPI(string(jsonValue), key, "harvest")
				}()

				if err := <-errChan; err != nil {
					log.Fatalf("error in hit besu api: %s", err)
				}

			case "bc_distribution":
				jsonValue, _ := json.Marshal(dataMap)
				errChan := make(chan error, 1)
				go func() {
					errChan <- w.HitBesuAPI(string(jsonValue), key, "distribution")
				}()

				if err := <-errChan; err != nil {
					log.Fatalf("error in hit besu api: %s", err)
				}

			case "bc_seller":
				jsonValue, _ := json.Marshal(dataMap)
				errChan := make(chan error, 1)
				go func() {
					errChan <- w.HitBesuAPI(string(jsonValue), key, "seller-box")
				}()

				if err := <-errChan; err != nil {
					log.Fatalf("error in hit besu api: %s", err)
				}
			case "bc_collector":
				jsonValue, _ := json.Marshal(dataMap)
				errChan := make(chan error, 1)
				go func() {
					errChan <- w.HitBesuAPI(string(jsonValue), key, "harvest-collector")
				}()

				if err := <-errChan; err != nil {
					log.Fatalf("error in hit besu api: %s", err)
				}
			case "bc_processor":
				jsonValue, _ := json.Marshal(dataMap)
				errChan := make(chan error, 1)
				go func() {
					errChan <- w.HitBesuAPI(string(jsonValue), key, "harvest-processor")
				}()

				if err := <-errChan; err != nil {
					log.Fatalf("error in hit besu api: %s", err)
				}
			case "email_verify":
				email := dataMap["email"].(string)
				token := dataMap["token"].(string)
				fmt.Println("Processing:", op, email)
				helper.SendEmailValidateEmail(email, token)
			case "reset_password":
				email := dataMap["email"].(string)
				token := dataMap["token"].(string)
				helper.SendEmailResetPassword(email, token)
			default:
				log.Println("Unknown op:", op)
			}
		}

		w.Redis.Del(ctx, key)
	}
}

func (w *Worker) StartFlushWorker(interval time.Duration) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.running {
		return
	}

	w.ticker = time.NewTicker(interval)
	w.quit = make(chan struct{})
	w.running = true

	go func() {
		for {
			select {
			case <-w.ticker.C:
				w.flushPendingItems()
			case <-w.quit:
				w.ticker.Stop()
				return
			}
		}
	}()
}

func (w *Worker) StopFlushWorker() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.running {
		return
	}

	close(w.quit)
	w.running = false
}
func (w *Worker) HitBesuAPI(value, key, endpoint string) error {
	ctx := context.Background()
	client := &http.Client{Timeout: 2 * time.Second}

	// Decode JSON fleksibel
	var parsed map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(value))
	decoder.UseNumber() // biar angka tetap json.Number, bukan float64
	if err := decoder.Decode(&parsed); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Re-marshal body untuk dikirim ke API blockchain
	body, err := json.Marshal(parsed)
	if err != nil {
		return fmt.Errorf("failed to re-marshal JSON: %w", err)
	}

	// Kirim request ke API blockchain
	req, err := http.NewRequest("POST", os.Getenv("BLOCKCHAIN_API")+"/blockchain/"+endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s API returned non-200 status: %d", endpoint, resp.StatusCode)
	}

	// Decode response dari blockchain
	var respData struct {
		TxBlock string `json:"tx_hash"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	var idUint uint
	switch v := parsed["id"].(type) {
	case json.Number:
		i64, _ := v.Int64()
		idUint = uint(i64)
	case float64:
		idUint = uint(v)
	case string:
		i64, _ := strconv.ParseInt(v, 10, 64)
		idUint = uint(i64)
	default:
		return fmt.Errorf("invalid id type: %T", v)
	}

	switch endpoint {
	case "harvest":
		return w.farm.UpdateBcBlockHarvest(ctx, idUint, respData.TxBlock)
	case "distribution":
		return w.distribution.UpdateBcBlockDistribution(ctx, idUint, respData.TxBlock)
	case "seller-box":
		return w.seller.UpdateBcBlockSellerBox(ctx, idUint, respData.TxBlock)
	case "harvest-collector":
		return w.collector.UpdateBcBlockCollector(ctx, idUint, respData.TxBlock)
	case "harvest-processor":
		return w.processor.UpdateBcBlockProcessor(ctx, idUint, respData.TxBlock)
	default:
		return errors.New("invalid endpoint")
	}
}
