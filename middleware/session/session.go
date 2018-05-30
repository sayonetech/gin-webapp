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
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

const maxAge int = 365 * 24 * 60 * 60

var client *redis.Client

type Store interface {
	// Set sets value to given key in session.
	Save(context *gin.Context) error
	// Get gets value by given key in session.
	Get(context *gin.Context, key string) string //Session to be renamed
	// Delete deletes a key from session.
	Delete(context *gin.Context, key string) error
	// ID returns current session ID.
	ID() string
	// Release releases session resource and save data to provider.
	Decode(context *gin.Context) error
	// ID returns current session ID.
	Encode() string
	// Check the session object expiry
	IsExpired() bool
}

type SessionStore struct {
	session Session
	cache   *redis.Client
}

func NewSessionStore() *SessionStore {
	sessionStore := &SessionStore{
		cache: client,
	}
	return sessionStore
}

//Session ... The Base session class
type Session struct {
	SessionKey  string
	SessionData string
	ExpireDate  time.Time //604800 7 days
}

func (store *SessionStore) Save(context *gin.Context) {

}

func (store *SessionStore) Get(context *gin.Context, key string) string {
	return store.session.SessionKey
}

func (store *SessionStore) ID() string {
	return store.session.SessionKey
}

func (store *SessionStore) Decode(context *gin.Context) {
	//TODO return SessionData Struct
}

func (store *SessionStore) Encode() {

}
func (store *SessionStore) IsExpired() bool {
	return false
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
	userData, err := msgpack.Marshal(user)
	if err != nil {
		panic(err)
	}
	encrypted, err := encrypt(config.GetSessionConfig().Secret, userData)
	if err != nil {

		log.WithFields(log.Fields{
			"user": fmt.Sprint(user.ID),
		}).Info("unable to encode", err)
		return false, &sessionError{"error with encrypting the session key. Check the session configuration"}

	}
	sessionToken := sessionId()
	session := Session{SessionKey: sessionToken, SessionData: encrypted}
	store := Default(context)
	store.session = session
	store.Save(context)
	//github.com/vmihailenco/msgpack
	//Set Cookie
	setSessionCookie(context, session)

	//Set ExpireDate
	//Create new session and save

	return true, nil
}

//SetSessionCookie ... Set Cookie after the authentication
//https://stackoverflow.com/questions/40887538/go-gin-unable-to-set-cookies
func setSessionCookie(context *gin.Context, session Session) {
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

// shortcut to get session
func Default(c *gin.Context) SessionStore {
	return c.MustGet("store").(SessionStore)
}

func init() {
	connection := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client = connection
}
