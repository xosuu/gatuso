package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)





func LoadDatabase()*gorm.DB{
	
	dsn := "host=localhost user=mikis password=mikis123 dbname=pelisdb port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), nil)
	if(err != nil){
		panic(err)
	}
	log.Println("Connection success")
	return DB

}


