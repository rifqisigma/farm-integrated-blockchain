package main

import (
	"farm-integrated-web3/cmd/database"
	"farm-integrated-web3/entity"
	"log"
)

func main() {

	db, _, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		// User dan token harus paling awal
		&entity.User{},
		&entity.Token{},

		// Geo data
		&entity.Country{},
		&entity.Region{},
		&entity.Regency{},

		// Master data
		&entity.Crop{},

		// Profile (karena child-nya nanti butuh ini)
		&entity.FarmerProfile{},
		&entity.DistributorProfile{},
		&entity.SellerProfile{},
		&entity.ProcessorProfile{},
		&entity.CollectorProfile{},
		&entity.ConsumerProfile{},

		// Transaction data
		&entity.Harvest{},
		&entity.HarvestCollector{},
		&entity.HarvestProcessor{},
		&entity.Distribution{},
		&entity.SellerBox{},
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrasi selesai dan database siap")
}
