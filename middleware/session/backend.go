package session

import (
	"go-webapp/models"

	"github.com/jinzhu/gorm"
)

//https://github.com/apexskier/httpauth/blob/master/auth.go

// The AuthBackend interface defines a set of methods an AuthBackend must implement.
type AuthBackend interface {
	SaveUser(u models.User) error
	User(id int) (user models.User, e error)
	DeleteUser(id int) error
}

//Backend Driver
type Backend struct {
	driverName string
	db         *gorm.DB
}

//NewBackend initilise the backend object
func NewBackend() (b Backend) {
	b.driverName = "postgres"
	b.db = models.Model
	return b
}

// User returns the user with the given usedID. Error is set to
// ErrMissingUser if user is not found.
func (b Backend) User(id int) (user models.User, e error) {
	return user, nil
}

// SaveUser adds a new user
func (b Backend) SaveUser(u models.User) error {
	return nil
}

// DeleteUser removes a user, raising ErrDeleteNull if that user was missing.
func (b Backend) DeleteUser(id int) error {
	return nil
}
