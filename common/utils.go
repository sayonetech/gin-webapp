package common

// Common tools and helper functions
import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"go-webapp/validators"
	validator "gopkg.in/go-playground/validator.v8"
)

type UserModelValidator struct {
	Username  string `validate:"required" form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	Email     string `form:"email" json:"email" binding:"exists,email"`
	Password  string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Phone     string `form:"phone" json:"phone"`
	//UserModel models.User `json:"-"`
}

const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(obj *UserModelValidator) error {
	//b := binding.Default(c.Request.Method, c.ContentType())
	//return c.ShouldBindWith(obj, b)
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	return validate.Struct(obj)

}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

//GetEnv ... to get the value from enviornment
func Getenv(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

// GetClientIP gets the correct IP for the end client instead of the proxy
func GetClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
