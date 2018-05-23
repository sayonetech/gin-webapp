package session

import "github.com/gin-gonic/gin"

const key = "session_store"

// FromContext returns the Store associated with this context.
func FromContext(context *gin.Context) Store {
	return context.MustGet(key).(Store)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(context *gin.Context, store Store) {
	context.Set(key, store)
}
