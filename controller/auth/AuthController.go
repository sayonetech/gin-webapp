package auth

import (
	"go-webapp/common"
	"go-webapp/models"
	"go-webapp/serializer"
	"go-webapp/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	//https://medium.com/@etiennerouzeaud/how-to-create-a-basic-restful-api-in-go-c8e032ba3181
	//https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/routers.go
	userModelValidator := validators.NewUserModelValidator()
	if err := userModelValidator.Bind(context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := models.SaveOne(&userModelValidator.UserModel); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	context.Set("user", userModelValidator.UserModel)
	serializer := serializer.UserSerializer{Context: context}
	context.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func Login(c *gin.Context) {
	//https://github.com/demo-apps/go-gin-app/blob/master/routes.go
}

func LogOut(c *gin.Context) {
	//https://github.com/demo-apps/go-gin-app/blob/master/routes.go
}
