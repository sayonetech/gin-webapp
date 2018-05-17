package session

import (
	"crypto/rand"
	"encoding/base64"
	"go-webapp/config"
	"go-webapp/models"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

const maxAge int = 365 * 24 * 60 * 60

type SessionData struct {
	UserID    int
	UserEmail string
}

type SessionStore interface {
	// Set sets value to given key in session.
	Set(interface{}, interface{}) error
	// Get gets value by given key in session.
	Get(interface{}) interface{}
	// Delete deletes a key from session.
	Delete(interface{}) error
	// ID returns current session ID.
	ID() string
	// Release releases session resource and save data to provider.
	Decode(interface{}) error
	// Flush deletes all session data.
	Flush() error
	// ID returns current session ID.
	Encode() string
}

//Session ... The Base session class
type Session struct {
	SessionKey  string
	SessionData SessionData
	ExpireDate  time.Time
}

func (session *Session) Decode() {
	//TODO return SessionData Struct
}

func (session *Session) Encode() {

}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//Authenticate ... Authenticate the user with session
func Authenticate(context *gin.Context, user models.User) {
	//Encode user data
	//Set ExpireDate
	//Create new session and save
	//Set Cookie
}

//SetSessionCookie ... Set Cookie after the authentication
//https://stackoverflow.com/questions/40887538/go-gin-unable-to-set-cookies
func SetSessionCookie(context *gin.Context) {
	//TODO Create Session Object
	//TODO Set the cookie
	//TODO Save Session to redis
	sessionToken := sessionId()

	context.SetCookie(config.GetSessionConfig().Name,
		sessionToken,
		config.GetSessionConfig().MaxAge,
		config.GetSessionConfig().Path,
		config.GetSessionConfig().Domain,
		config.GetSessionConfig().Secure,
		config.GetSessionConfig().HttpOnly,
	)

}
