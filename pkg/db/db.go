package db

import (
	"flag"
	"log"

	"github.com/heydp/oms/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbCmd *flag.FlagSet) *gorm.DB {
	dbURL := "postgres://postgres:@localhost:5432/ordertry"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	// db, err := gorm.Open("postgres", "user=postgres password= dbname=order sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}
	db.Migrator().DropTable("all_items", "orders")
	err = db.AutoMigrate(&models.Order{}, &models.AllItem{})
	if err != nil {
		log.Fatalln("Err in migration", err)
	}

	return db
}

func InitTest() error {
	dbURL := "postgres://postgres:@localhost:5432/order"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.AllItem{})
	log.Println("Successfully connected to db ", db)
	return nil
}
