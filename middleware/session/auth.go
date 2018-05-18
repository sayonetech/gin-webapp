package session

// Authorizer structures contain the store of user session cookies a reference
// to a backend storage system.
type Authorizer struct {
	session string      //redis session TODO
	backend AuthBackend //auth backend TODO
}

//todo https://github.com/apexskier/httpauth/blob/master/auth.go
