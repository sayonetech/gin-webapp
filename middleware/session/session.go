package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go-webapp/config"
	"go-webapp/models"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const maxAge int = 365 * 24 * 60 * 60

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
	// Check the session object expiry
	IsExpired() bool
}

//Session ... The Base session class
type Session struct {
	SessionKey  string
	SessionData string
	ExpireDate  time.Time //604800 7 days
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
func Authenticate(context *gin.Context, user models.User) (bool, error) {
	//Encode user data
	encrypted, err := encrypt(config.GetSessionConfig().Secret, fmt.Sprint(user.ID))
	if err != nil {

		log.WithFields(log.Fields{
			"user": fmt.Sprint(user.ID),
		}).Info("unable to encode", err)
		return false, &sessionError{"error with encrypting the session key. Check the session configuration"}

	}
	sessionToken := sessionId()
	session := &Session{SessionKey: sessionToken, SessionData: encrypted}
	//Set Cookie
	setSessionCookie(context, session)

	//Set ExpireDate
	//Create new session and save

	return true, nil
}

//SetSessionCookie ... Set Cookie after the authentication
//https://stackoverflow.com/questions/40887538/go-gin-unable-to-set-cookies
func setSessionCookie(context *gin.Context, session *Session) {
	//TODO Create Session Object
	//TODO Set the cookie
	//TODO Save Session to redis

	context.SetCookie(config.GetSessionConfig().Name,
		session.SessionKey,
		config.GetSessionConfig().MaxAge,
		config.GetSessionConfig().Path,
		config.GetSessionConfig().Domain,
		config.GetSessionConfig().Secure,
		config.GetSessionConfig().HttpOnly,
	)

}
