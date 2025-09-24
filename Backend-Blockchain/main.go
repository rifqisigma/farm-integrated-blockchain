package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	messagestore "besu-go/message"
)

var (
	client      *ethclient.Client
	auth        *bind.TransactOpts
	contract    *messagestore.Messagestore
	contractAdr common.Address
)

func init() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ no .env file found, using system env")
	}

	// Connect ke Besu
	rpcURL := os.Getenv("BESU_URL")
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("❌ gagal konek ke Besu: ", err)
	}

	// Private Key
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("❌ invalid private key: ", err)
	}

	// Auth (Transactor)
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // chainID custom, samain dgn config Besu
	if err != nil {
		log.Fatal("❌ failed make transactor: ", err)
	}

	// Address contract yg sudah di-deploy
	contractAdr = common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	// Instance contract
	contract, err = messagestore.NewMessagestore(contractAdr, client)
	if err != nil {
		log.Fatal("❌ gagal load contract: ", err)
	}

	fmt.Println("✅ Connected to contract at", contractAdr.Hex())
}

func main() {
	app := fiber.New()

	// Insert data
	// Insert data
	app.Post("/insert", func(c *fiber.Ctx) error {
		var req struct {
			ID      int64  `json:"id"`
			Status  string `json:"status"`
			Message string `json:"message"`
			To      string `json:"to"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		toAddr := common.HexToAddress(req.To)

		tx, err := contract.CreateMessage(auth,
			big.NewInt(req.ID),
			req.Status,
			req.Message,
			toAddr,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"tx_hash": tx.Hash().Hex(),
		})
	})

	// Get data by ID
	// Get data by index (ID)
	app.Get("/get/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		idBig := new(big.Int)
		idBig, ok := idBig.SetString(idParam, 10)
		if !ok {
			return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
		}

		resID, status, message, from, to, createdAt, err := contract.GetMessage(&bind.CallOpts{Context: context.Background()}, idBig)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"id":         resID.String(),
			"status":     status,
			"message":    message,
			"from":       from.Hex(),
			"to":         to.Hex(),
			"created_at": createdAt.String(),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
