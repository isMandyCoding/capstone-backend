package databaseConfig

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func dbStart(*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=volunteer-connect password=12345")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	database := db.DB()

	err = database.Ping()

	if err != nil {
		panic(err.Error())
	}
}
