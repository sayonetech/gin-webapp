package auth

import (
	"go-webapp/common"
	"go-webapp/middleware/auth"
	"go-webapp/serializer"
	"go-webapp/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	backend    auth.Backend
	Authorizer *auth.Authorizer
)

func init() {
	backend = auth.NewBackend()
	Authorizer = auth.NewAuthorizer(backend)
}

func Register(context *gin.Context) {
	//https://medium.com/@etiennerouzeaud/how-to-create-a-basic-restful-api-in-go-c8e032ba3181
	//https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/routers.go
	userModelValidator := validators.NewUserModelValidator()

	if err := userModelValidator.Bind(context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := Authorizer.Register(&userModelValidator.UserModel); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	context.Set("user", userModelValidator.UserModel)
	serializer := serializer.UserSerializer{Context: context}
	context.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UserLogin(context *gin.Context) {
	loginValidator := validators.NewLoginValidator()
	if err := loginValidator.Bind(context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	Authorizer.Login(context, loginValidator.Email, loginValidator.Password)
	//https://github.com/acoshift/session
	//https://github.com/go-macaron/session
	//https://github.com/knq/sessionmw
	//https://github.com/gin-contrib/sessions
	//https://stackoverflow.com/questions/47085046/gin-sessions-stores-the-status-and-the-code-in-the-url-i-want-to-change-that-t
	//https://www.sohamkamani.com/blog/2018/03/25/golang-session-authentication/
	//https://github.com/apexskier/httpauth
	//https://github.com/rageix/ginAuth
	//https://jonathanmh.com/go-gin-http-basic-auth/

}

func LogOut(c *gin.Context) {
	//https://github.com/demo-apps/go-gin-app/blob/master/routes.go
}
