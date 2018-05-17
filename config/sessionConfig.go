package config

// Config contains the configuration parameters for the session middleware.
type Config struct {
	// Secret is
	Secret string

	// Name is the cookie name.
	Name string

	// Path is the cookie path.
	Path string

	// Domain is the cookie domain.
	Domain string

	// MaxAge is the cookie max age.
	MaxAge int

	// Secure is the cookie secure flag.
	Secure bool

	// HttpOnly is the cookie http only flag.
	HttpOnly bool
}
