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
		&entity.User{},
		&entity.ConsumerProfile{},
		&entity.DistributorProfile{},
		&entity.FarmerProfile{},
		&entity.RetailerProfile{},
		&entity.Crop{},
		&entity.Harvest{},
		&entity.Distribution{},
		&entity.RetailerCart{},
		&entity.Token{},
		&entity.Country{},
		&entity.Region{},
		&entity.Regency{},
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrasi selesai dan database siap")
}
