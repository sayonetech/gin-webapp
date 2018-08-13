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

type Store struct {
	session Session
	cache   *redis.Client
}

func NewSessionStore() *Store {
	store := &Store{
		cache: client,
	}
	return store
}

//Session ... The Base session class
type Session struct {
	SessionKey  string
	SessionData string
}

func (store *Store) Save(context *gin.Context) {
	err := store.cache.Set(store.session.SessionKey, store.session.SessionData, time.Hour).Err()
	if err != nil {
		panic(err)
	}
	log.WithFields(log.Fields{
		"test": "ddddd",
	}).Info("Save SessionStore")
}

func (store *Store) Authenticate(context *gin.Context, user models.User) bool {
	encrypted, err := store.Encode(context, user)
	if err != nil {
		log.WithFields(log.Fields{
			"authenticate": "Error occured",
		}).Info("unable to encode", err)
		return false
	}
	sessionToken := sessionId()
	session := Session{SessionKey: sessionToken, SessionData: encrypted}
	store.session = session
	store.Save(context)
	setSessionCookie(context, session)
	return true
}

func (store *Store) Get(context *gin.Context, key string) string {
	return store.session.SessionKey
}

func (store *Store) ID() string {
	return store.session.SessionKey
}

func (store *Store) Encode(context *gin.Context, user models.User) (string, error) {
	userData, err := msgpack.Marshal(user)
	if err != nil {
		panic(err)
	}

	encrypted, err := encrypt(config.GetSessionConfig().Secret, userData)
	if err != nil {

		log.WithFields(log.Fields{
			"user": fmt.Sprint(user.ID),
		}).Info("unable to encode", err)
		return "", &sessionError{"error with encrypting the session key. Check the session configuration"}

	}
	return encrypted, nil
}

func (store *Store) Decode() {

}
func (store *Store) IsExpired() bool {
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
func Authenticate(context *gin.Context, user models.User) bool {
	store := Default(context)
	return store.Authenticate(context, user)
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
func Default(c *gin.Context) *Store {
	return c.MustGet("store").(*Store)
}

func init() {
	connection := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client = connection
}
