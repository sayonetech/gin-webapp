package session

import "beco/models"

//https://github.com/apexskier/httpauth/blob/master/auth.go

// The AuthBackend interface defines a set of methods an AuthBackend must
// implement.
type AuthBackend interface {
	SaveUser(u models.User) error
	User(id int) (user models.User, e error)
	DeleteUser(id int) error
	Close()
}
