package serializer

import (
	"go-webapp/models"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	Context *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (self *UserSerializer) Response() UserResponse {
	userModel := self.Context.MustGet("user").(models.User)
	user := UserResponse{
		Username: userModel.UserName,
		Email:    userModel.Email,
	}
	return user
}
