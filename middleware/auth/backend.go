package auth

import (
	"go-webapp/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//https://github.com/apexskier/httpauth/blob/master/auth.go

// The AuthBackend interface defines a set of methods an AuthBackend must implement.
type AuthBackend interface {
	SaveUser(data interface{}) error
	FetchUser(condition interface{}) (user models.User, e error)
	DeleteUser(condition interface{}) error
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
// Error if user is not found.
func (b Backend) FetchUser(condition interface{}) (user models.User, e error) {
	var model models.User
	err := b.db.Where(condition).First(&model).Error
	return model, err
}

// SaveUser adds a new user
func (b Backend) SaveUser(data interface{}) error {
	err := b.db.Save(data).Error
	log.WithFields(log.Fields{
		"error": err,
	}).Info("Backend")
	return err
}

// DeleteUser removes a user, raising ErrDeleteNull if that user was missing.
func (b Backend) DeleteUser(condition interface{}) error {
	return nil
}
