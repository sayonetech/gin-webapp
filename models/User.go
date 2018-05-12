package models

import (
	"time"

	"github.com/jinzhu/gorm"
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
