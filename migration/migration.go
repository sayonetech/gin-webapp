package main

import (
	"fmt"
	"go-webapp/config"
	m "go-webapp/models"

	"github.com/jinzhu/gorm"
)

func main() {

	database, err := gorm.Open("postgres", "host="+config.GetEnv().HOST+" port="+config.GetEnv().DATABASE_PORT+" user="+config.GetEnv().DATABASE_USERNAME+" dbname="+config.GetEnv().DATABASE_NAME+" password="+config.GetEnv().DATABASE_PASSWORD+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Applying migration")
	fmt.Println("Applying user migration")
	database.AutoMigrate(&m.User{})
	fmt.Println("Finished migration")
}
