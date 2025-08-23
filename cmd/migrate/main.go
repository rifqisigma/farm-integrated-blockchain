package migrate

import (
	"farm-integrated-web3/cmd/database"
	"log"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migrasi selesai dan database siap")
}
