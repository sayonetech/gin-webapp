package models

import (
	"go-webapp/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Model *gorm.DB

// Read more about init() https://stackoverflow.com/questions/17733220/go-package-initialization
func init() {
	var err error

	Model, err = gorm.Open("postgres", "host="+config.GetEnv().HOST+" port="+config.GetEnv().DATABASE_PORT+" user="+config.GetEnv().DATABASE_USERNAME+" dbname="+config.GetEnv().DATABASE_NAME+" password="+config.GetEnv().DATABASE_PASSWORD+" sslmode=disable")

	if err != nil {
		panic(err)
	}
}
