package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User ...User Model
type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	UserName  string
	FirstName string
	LastName  string
	IsActive  bool
	Password  string
	Phone     string
	LastLogin time.Time
}

func (User) TableName() string {
	return "users"
}

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

//Refer https://github.com/demo-apps/go-gin-app
//https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/models.go
//https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
//https://github.com/gin-contrib
//https://github.com/vsouza/go-gin-boilerplate/tree/master/middlewares\
//https://github.com/george518/PPGo_Api_Demo_Gin/tree/master/routers
//https://github.com/gothinkster/golang-gin-realworld-example-app
// swagger & user https://github.com/EDDYCJY/go-gin-example
//https://github.com/night-codes/summer
//https://github.com/szuecs/gin-gomonitor
//https://github.com/sbecker/gin-api-demo
//https://github.com/acrosson/gingorm/tree/master/db
//https://github.com/aviddiviner/gin-limit
//https://github.com/Luncher/go-rest
//https://github.com/nightlegend/apigateway
//https://github.com/rageix/ginAuth
