package auth

import (
	"errors"
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

func UserLogin(context *gin.Context) {
	loginValidator := validators.NewLoginValidator()
	if err := loginValidator.Bind(context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := models.FindUser(&models.User{Email: loginValidator.Email})

	if err != nil {
		context.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.CheckPassword(loginValidator.Password) != nil {
		context.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}
	//Manage session here
	context.JSON(http.StatusOK, gin.H{})
}

func LogOut(c *gin.Context) {
	//https://github.com/demo-apps/go-gin-app/blob/master/routes.go
}
