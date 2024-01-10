package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDataBase() *Database {
	return &Database{}
}
func (db *Database) ConnectToDb() {
	var err error
	var dataBaseString string

	dataBaseString = fmt.Sprintf("host=admin user=admin password=admin dbname=databasename  port=90")
	db.DB,err = gorm.Open(postgres.Open(dataBaseString), &gorm.Config{})
	if err != nil {
		log.Fatal("fall connect to db")
	}
}