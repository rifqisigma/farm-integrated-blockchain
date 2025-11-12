package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	supplychain "besu-go/supplychain" // import hasil abigen Go
)

var (
	client      *ethclient.Client
	auth        *bind.TransactOpts
	contract    *supplychain.Contract
	contractAdr common.Address
	txHash      string
)

// updateOrAddEnv menulis atau menambahkan key=value di file .env
func updateOrAddEnv(filePath, key, value string) error {
	lines := []string{}
	found := false

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = fmt.Sprintf("%s=%s", key, value)
			found = true
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func main() {
	fmt.Println("MAIN API - Connecting to Blockchain...")

	// Load env kalau ada
	_ = godotenv.Load()

	rpcURL := os.Getenv("BESU_URL")
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	existingAddress := os.Getenv("CONTRACT_ADDRESS")
	existingTx := os.Getenv("TX_HASH")

	// Connect ke Besu
	var err error
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("❌ Gagal konek ke Besu:", err)
	}

	// Private key & auth
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("❌ Invalid private key:", err)
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatal("❌ Gagal membuat transactor:", err)
	}
	auth.GasLimit = 3_000_000

	// Cek contract, kalau kosong deploy baru
	if existingAddress == "" || existingAddress == "0x0000000000000000000000000000000000000000" {
		fmt.Println("🚀 Deploying smart contract (first time)...")
		addr, tx, _, err := supplychain.DeployContract(auth, client)
		if err != nil {
			log.Fatal("❌ Gagal deploy kontrak:", err)
		}

		// Update variabel lokal & env
		contractAdr = addr
		txHash = tx.Hash().Hex()
		if err := updateOrAddEnv(".env", "CONTRACT_ADDRESS", contractAdr.Hex()); err != nil {
			log.Fatal("❌ Gagal write env contract address:", err)
		}
		if err := updateOrAddEnv(".env", "TX_HASH", txHash); err != nil {
			log.Fatal("❌ Gagal write env tx hash:", err)
		}

		fmt.Printf("🚀 Contract deployed at: %s\n📦 TX Hash: %s\n✅ Saved to .env\n", contractAdr.Hex(), txHash)

	} else {
		// Pakai yang ada di env
		contractAdr = common.HexToAddress(existingAddress)
		txHash = existingTx
		fmt.Printf("ℹ️  Using existing contract at: %s\n", contractAdr.Hex())
	}

	// Load contract instance
	contract, err = supplychain.NewContract(contractAdr, client)
	if err != nil {
		log.Fatal("❌ Gagal load contract:", err)
	}

	fmt.Println("✅ Connected to contract at", contractAdr.Hex())

	app := fiber.New()
	bc := app.Group("/blockchain")

	// ====================== HARVEST ======================
	bc.Post("/harvest", createHarvest)
	bc.Get("/harvest/:id", getHarvest)
	bc.Get("/harvest", getAllHarvest)

	// ====================== HARVEST COLLECTOR ======================
	bc.Post("/harvest-collector", createHarvestCollector)
	bc.Get("/harvest-collector/:id", getHarvestCollector)
	bc.Get("/harvest-collector", getAllHarvestCollector)

	// ====================== HARVEST PROCESSOR ======================
	bc.Post("/harvest-processor", createHarvestProcessor)
	bc.Get("/harvest-processor/:id", getHarvestProcessor)
	bc.Get("/harvest-processor", getAllHarvestProcessor)

	// ====================== DISTRIBUTION ======================
	bc.Post("/distribution", createDistribution)
	bc.Get("/distribution/:id", getDistribution)
	bc.Get("/distribution", getAllDistribution)

	// ====================== SELLER BOX ======================
	bc.Post("/seller-box", createSellerBox)
	bc.Get("/seller-box/:id", getSellerBox)
	bc.Get("/seller-box", getAllSellerBox)

	// Jalankan server
	log.Fatal(app.Listen(":8001"))
}

// ====================== HARVEST ======================
func createHarvest(c *fiber.Ctx) error {
	var req struct {
		ID          int64  `json:"id"`
		FarmerID    int64  `json:"farmer_profile_id"`
		CropID      int64  `json:"crop_id"`
		RegencyID   int64  `json:"regency_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Quantity    int64  `json:"quantity"`
		BasePrice   int64  `json:"base_price"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	tx, err := contract.CreateHarvest(auth,
		big.NewInt(req.ID),
		big.NewInt(req.FarmerID),
		big.NewInt(req.CropID),
		big.NewInt(req.RegencyID),
		req.Name,
		req.Description,
		big.NewInt(req.Quantity),
		big.NewInt(req.BasePrice),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"tx_hash": tx.Hash().Hex()})
}

func getHarvest(c *fiber.Ctx) error {
	idBig, ok := new(big.Int).SetString(c.Params("id"), 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	// Cek apakah ID 0 (data tidak ada)
	if idBig.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest not found"})
	}

	res, err := contract.Harvests(&bind.CallOpts{Context: context.Background()}, idBig)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch harvest", "detail": err.Error()})
	}

	// Cek apakah data yang di-return kosong (ID = 0 artinya tidak ada)
	if res.Id.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest not found"})
	}

	// Format response
	harvest := fiber.Map{
		"id":                res.Id.Int64(),
		"farmer_profile_id": res.FarmerProfileId.Int64(),
		"crop_id":           res.CropId.Int64(),
		"regency_id":        res.RegencyId.Int64(),
		"name":              res.Name,
		"description":       res.Description,
		"quantity":          res.Quantity.Int64(),
		"base_price":        res.BasePrice.Int64(),
		"created_at":        res.CreatedAt.Int64(),
	}
	return c.JSON(harvest)
}

func getAllHarvest(c *fiber.Ctx) error {
	// Panggil fungsi yang benar dari Solidity: getAllHarvestTuples
	// Return: ids, farmerIds, cropIds, regencyIds, names, descs, quantities, basePrices, createdAts
	ids, farmerIds, cropIds, regencyIds, names, descs, quantities, basePrices, createdAts, err := contract.GetAllHarvestTuples(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse tuple menjadi array of objects
	var harvests []fiber.Map
	for i := 0; i < len(ids); i++ {
		harvest := fiber.Map{
			"id":                ids[i].Int64(),
			"farmer_profile_id": farmerIds[i].Int64(),
			"crop_id":           cropIds[i].Int64(),
			"regency_id":        regencyIds[i].Int64(),
			"name":              names[i],
			"description":       descs[i],
			"quantity":          quantities[i].Int64(),
			"base_price":        basePrices[i].Int64(),
			"created_at":        createdAts[i].Int64(),
		}
		harvests = append(harvests, harvest)
	}
	return c.JSON(harvests)
}

// ====================== HARVEST COLLECTOR ======================
func createHarvestCollector(c *fiber.Ctx) error {
	var req struct {
		ID          int64  `json:"id"`
		CollectorID int64  `json:"collector_profile_id"`
		HarvestID   int64  `json:"harvest_id"`
		Name        string `json:"name"`
		Desc        string `json:"desc"`
		Quantity    int64  `json:"quantity"`
		Price       int64  `json:"price"`
		BasePrice   int64  `json:"base_price"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	tx, err := contract.CreateHarvestCollector(auth,
		big.NewInt(req.ID),
		big.NewInt(req.CollectorID),
		big.NewInt(req.HarvestID),
		req.Name,
		req.Desc,
		big.NewInt(req.Quantity),
		big.NewInt(req.Price),
		big.NewInt(req.BasePrice),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"tx_hash": tx.Hash().Hex()})
}

func getHarvestCollector(c *fiber.Ctx) error {
	idBig, ok := new(big.Int).SetString(c.Params("id"), 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if idBig.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest collector not found"})
	}

	res, err := contract.HarvestCollectors(&bind.CallOpts{Context: context.Background()}, idBig)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch harvest collector", "detail": err.Error()})
	}

	if res.Id.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest collector not found"})
	}

	// Format response
	harvestCollector := fiber.Map{
		"id":                   res.Id.Int64(),
		"collector_profile_id": res.CollectorProfileId.Int64(),
		"harvest_id":           res.HarvestId.Int64(),
		"name":                 res.Name,
		"desc":                 res.Desc,
		"quantity":             res.Quantity.Int64(),
		"price":                res.Price.Int64(),
		"base_price":           res.BasePrice.Int64(),
		"created_at":           res.CreatedAt.Int64(),
	}
	return c.JSON(harvestCollector)
}

func getAllHarvestCollector(c *fiber.Ctx) error {
	// Panggil fungsi yang benar dari Solidity: getAllHarvestCollectorTuples
	// Return: ids, collectorIDs, harvestIDs, names, descs, quantities, prices, basePrices
	ids, collectorIDs, harvestIDs, names, descs, quantities, prices, basePrices, err := contract.GetAllHarvestCollectorTuples(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse tuple menjadi array of objects
	var harvestCollectors []fiber.Map
	for i := 0; i < len(ids); i++ {
		hc := fiber.Map{
			"id":                   ids[i].Int64(),
			"collector_profile_id": collectorIDs[i].Int64(),
			"harvest_id":           harvestIDs[i].Int64(),
			"name":                 names[i],
			"desc":                 descs[i],
			"quantity":             quantities[i].Int64(),
			"price":                prices[i].Int64(),
			"base_price":           basePrices[i].Int64(),
		}
		harvestCollectors = append(harvestCollectors, hc)
	}
	return c.JSON(harvestCollectors)
}

// ====================== HARVEST PROCESSOR ======================
func createHarvestProcessor(c *fiber.Ctx) error {
	var req struct {
		ID                 int64  `json:"id"`
		ProcessorID        int64  `json:"processor_profile_id"`
		HarvestCollectorID int64  `json:"harvest_collector_id"`
		HarvestID          int64  `json:"harvest_id"`
		Name               string `json:"name"`
		Desc               string `json:"desc"`
		Quantity           int64  `json:"quantity"`
		BasePrice          int64  `json:"base_price"`
		Price              int64  `json:"price"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	tx, err := contract.CreateHarvestProcessor(auth,
		big.NewInt(req.ID),
		big.NewInt(req.ProcessorID),
		big.NewInt(req.HarvestCollectorID),
		big.NewInt(req.HarvestID),
		req.Name,
		req.Desc,
		big.NewInt(req.Quantity),
		big.NewInt(req.BasePrice),
		big.NewInt(req.Price),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"tx_hash": tx.Hash().Hex()})
}

func getHarvestProcessor(c *fiber.Ctx) error {
	idBig, ok := new(big.Int).SetString(c.Params("id"), 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if idBig.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest processor not found"})
	}

	res, err := contract.HarvestProcessors(&bind.CallOpts{Context: context.Background()}, idBig)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch harvest processor", "detail": err.Error()})
	}

	if res.Id.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "harvest processor not found"})
	}

	// Format response
	harvestProcessor := fiber.Map{
		"id":                   res.Id.Int64(),
		"processor_profile_id": res.ProcessorProfileId.Int64(),
		"harvest_collector_id": res.HarvestCollectorId.Int64(),
		"harvest_id":           res.HarvestId.Int64(),
		"name":                 res.Name,
		"desc":                 res.Desc,
		"quantity":             res.Quantity.Int64(),
		"base_price":           res.BasePrice.Int64(),
		"price":                res.Price.Int64(),
		"created_at":           res.CreatedAt.Int64(),
	}
	return c.JSON(harvestProcessor)
}

func getAllHarvestProcessor(c *fiber.Ctx) error {
	// Panggil fungsi yang benar dari Solidity: getAllHarvestProcessorTuples
	// Return: ids, processorIDs, collectorIDs, harvestIDs, names, descs, quantities, basePrices, prices
	ids, processorIDs, collectorIDs, harvestIDs, names, descs, quantities, basePrices, prices, err := contract.GetAllHarvestProcessorTuples(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse tuple menjadi array of objects
	var harvestProcessors []fiber.Map
	for i := 0; i < len(ids); i++ {
		hp := fiber.Map{
			"id":                   ids[i].Int64(),
			"processor_profile_id": processorIDs[i].Int64(),
			"harvest_collector_id": collectorIDs[i].Int64(),
			"harvest_id":           harvestIDs[i].Int64(),
			"name":                 names[i],
			"desc":                 descs[i],
			"quantity":             quantities[i].Int64(),
			"base_price":           basePrices[i].Int64(),
			"price":                prices[i].Int64(),
		}
		harvestProcessors = append(harvestProcessors, hp)
	}
	return c.JSON(harvestProcessors)
}

// ====================== DISTRIBUTION ======================
func createDistribution(c *fiber.Ctx) error {
	var req struct {
		ID               int64  `json:"id"`
		DistributorID    int64  `json:"distributor_profile_id"`
		DestinationID    int64  `json:"destination_id"`
		HarvestID        int64  `json:"harvest_id"`
		HarvestCollector int64  `json:"harvest_collector_id"`
		HarvestProcessor int64  `json:"harvest_processor_id"`
		Name             string `json:"name"`
		Desc             string `json:"desc"`
		Quantity         int64  `json:"quantity"`
		BasePrice        int64  `json:"base_price"`
		Price            int64  `json:"price"`
		Transportation   string `json:"transportation"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	tx, err := contract.CreateDistribution(auth,
		big.NewInt(req.ID),
		big.NewInt(req.DistributorID),
		big.NewInt(req.DestinationID),
		big.NewInt(req.HarvestID),
		big.NewInt(req.HarvestCollector),
		big.NewInt(req.HarvestProcessor),
		req.Name,
		req.Desc,
		big.NewInt(req.Quantity),
		big.NewInt(req.BasePrice),
		big.NewInt(req.Price),
		req.Transportation,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"tx_hash": tx.Hash().Hex()})
}

func getDistribution(c *fiber.Ctx) error {
	idBig, ok := new(big.Int).SetString(c.Params("id"), 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if idBig.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "distribution not found"})
	}

	res, err := contract.Distributions(&bind.CallOpts{Context: context.Background()}, idBig)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch distribution", "detail": err.Error()})
	}

	if res.Id.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "distribution not found"})
	}

	// Format response
	distribution := fiber.Map{
		"id":                     res.Id.Int64(),
		"distributor_profile_id": res.DistributorProfileId.Int64(),
		"destination_id":         res.DestinationId.Int64(),
		"harvest_id":             res.HarvestId.Int64(),
		"harvest_collector_id":   res.HarvestCollectorId.Int64(),
		"harvest_processor_id":   res.HarvestProcessorId.Int64(),
		"name":                   res.Name,
		"desc":                   res.Desc,
		"quantity":               res.Quantity.Int64(),
		"base_price":             res.BasePrice.Int64(),
		"price":                  res.Price.Int64(),
		"transportation":         res.Transportation,
		"created_at":             res.CreatedAt.Int64(),
	}
	return c.JSON(distribution)
}

func getAllDistribution(c *fiber.Ctx) error {
	// Panggil fungsi yang benar dari Solidity: getAllDistributionTuples
	// Return: ids, distributorIDs, destinationIDs, harvestIDs, collectorIDs, processorIDs, names, descs, quantities, basePrices, prices, transportations
	ids, distributorIDs, destinationIDs, harvestIDs, collectorIDs, processorIDs, names, descs, quantities, basePrices, prices, transportations, err := contract.GetAllDistributionTuples(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse tuple menjadi array of objects
	var distributions []fiber.Map
	for i := 0; i < len(ids); i++ {
		dist := fiber.Map{
			"id":                     ids[i].Int64(),
			"distributor_profile_id": distributorIDs[i].Int64(),
			"destination_id":         destinationIDs[i].Int64(),
			"harvest_id":             harvestIDs[i].Int64(),
			"harvest_collector_id":   collectorIDs[i].Int64(),
			"harvest_processor_id":   processorIDs[i].Int64(),
			"name":                   names[i],
			"desc":                   descs[i],
			"quantity":               quantities[i].Int64(),
			"base_price":             basePrices[i].Int64(),
			"price":                  prices[i].Int64(),
			"transportation":         transportations[i],
		}
		distributions = append(distributions, dist)
	}
	return c.JSON(distributions)
}

// ====================== SELLER BOX ======================
func createSellerBox(c *fiber.Ctx) error {
	var req struct {
		ID             int64  `json:"id"`
		SellerID       int64  `json:"seller_profile_id"`
		DistributionID int64  `json:"distribution_id"`
		Name           string `json:"name"`
		Desc           string `json:"desc"`
		Quantity       int64  `json:"quantity"`
		BasePrice      int64  `json:"base_price"`
		Price          int64  `json:"price"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	tx, err := contract.CreateSellerBox(auth,
		big.NewInt(req.ID),
		big.NewInt(req.SellerID),
		big.NewInt(req.DistributionID),
		req.Name,
		req.Desc,
		big.NewInt(req.Quantity),
		big.NewInt(req.BasePrice),
		big.NewInt(req.Price),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"tx_hash": tx.Hash().Hex()})
}

func getSellerBox(c *fiber.Ctx) error {
	idBig, ok := new(big.Int).SetString(c.Params("id"), 10)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if idBig.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "seller box not found"})
	}

	res, err := contract.SellerBoxes(&bind.CallOpts{Context: context.Background()}, idBig)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch seller box", "detail": err.Error()})
	}

	if res.Id.Cmp(big.NewInt(0)) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "seller box not found"})
	}

	// Format response
	sellerBox := fiber.Map{
		"id":                res.Id.Int64(),
		"seller_profile_id": res.SellerProfileId.Int64(),
		"distribution_id":   res.DistributionId.Int64(),
		"name":              res.Name,
		"desc":              res.Desc,
		"quantity":          res.Quantity.Int64(),
		"base_price":        res.BasePrice.Int64(),
		"price":             res.Price.Int64(),
		"created_at":        res.CreatedAt.Int64(),
	}
	return c.JSON(sellerBox)
}

func getAllSellerBox(c *fiber.Ctx) error {
	// Panggil fungsi yang benar dari Solidity: getAllSellerBoxTuples
	// Return: ids, sellerIDs, distributionIDs, names, descs, quantities, basePrices, prices
	ids, sellerIDs, distributionIDs, names, descs, quantities, basePrices, prices, err := contract.GetAllSellerBoxTuples(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse tuple menjadi array of objects
	var sellerBoxes []fiber.Map
	for i := 0; i < len(ids); i++ {
		sb := fiber.Map{
			"id":                ids[i].Int64(),
			"seller_profile_id": sellerIDs[i].Int64(),
			"distribution_id":   distributionIDs[i].Int64(),
			"name":              names[i],
			"desc":              descs[i],
			"quantity":          quantities[i].Int64(),
			"base_price":        basePrices[i].Int64(),
			"price":             prices[i].Int64(),
		}
		sellerBoxes = append(sellerBoxes, sb)
	}
	return c.JSON(sellerBoxes)
}
