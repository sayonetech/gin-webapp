package main

import (
	"fmt"
	m "go-webapp/models"
)

func main() {
	fmt.Println("Applying migration")
	fmt.Println("Applying user migration")
	m.Model.AutoMigrate(&m.User{})
	fmt.Println("Finished migration")
}
