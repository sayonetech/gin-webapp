package validators

import (
	"go-webapp/common"
	"go-webapp/models"

	"github.com/gin-gonic/gin"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type UserModelValidator struct {
	User struct {
		Username  string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email     string `form:"email" json:"email" binding:"exists,email"`
		Password  string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		FirstName string `form:"first_name" json:"first_name" binding:"min=8,max=255"`
		LastName  string `form:"last_name" json:"image" binding:"min=8,max=255"`
		Phone     string `form:"phone" json:"image" binding:"min=10,max=10"`
	} `json:"user"`
	userModel models.User `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.UserName = self.User.Username
	self.userModel.Email = self.User.Email
	self.userModel.FirstName = self.User.FirstName
	self.userModel.LastName = self.User.LastName
	self.userModel.Phone = self.User.Phone

	if self.User.Password != common.NBRandomPassword {
		self.userModel.SetPassword(self.User.Password)
	}

	return nil
}

// You can put the default value of a Validator here
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel models.User) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.User.Username = userModel.UserName
	userModelValidator.User.Email = userModel.Email
	userModelValidator.User.FirstName = userModel.FirstName
	userModelValidator.User.LastName = userModel.LastName
	userModelValidator.User.Phone = userModel.Phone

	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password"json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
	userModel models.User `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.userModel.Email = self.User.Email
	return nil
}

// You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
