package validators

import (
	"go-webapp/common"
	"go-webapp/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
// type UserModelValidator struct {
// 	Username  string      `validate:"required" form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
// 	Email     string      `form:"email" json:"email" binding:"exists,email"`
// 	Password  string      `form:"password" json:"password" binding:"exists,min=8,max=255"`
// 	FirstName string      `form:"first_name" json:"first_name"`
// 	LastName  string      `form:"last_name" json:"last_name"`
// 	Phone     string      `form:"phone" json:"phone"`
// 	UserModel models.User `json:"-"`
// }

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *common.UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(self)

	if err != nil {
		log.WithFields(log.Fields{
			"self": self,
		}).Info("Error parsing", err)

		return err
	}
	self.UserModel.UserName = self.Username
	self.UserModel.Email = self.Email
	self.UserModel.FirstName = self.FirstName
	self.UserModel.LastName = self.LastName
	self.UserModel.Phone = self.Phone

	if self.Password != common.NBRandomPassword {
		self.UserModel.SetPassword(self.Password)
	}

	return nil
}

// You can put the default value of a Validator here
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel models.User) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.Username = userModel.UserName
	userModelValidator.Email = userModel.Email
	userModelValidator.FirstName = userModel.FirstName
	userModelValidator.LastName = userModel.LastName
	userModelValidator.Phone = userModel.Phone

	return userModelValidator
}

type LoginValidator struct {
	Email     string      `form:"email" json:"email" binding:"exists,email"`
	Password  string      `form:"password"json:"password" binding:"exists,min=8,max=255"`
	UserModel models.User `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.UserModel.Email = self.Email
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
