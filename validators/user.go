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
type UserModelValidator struct {
	User struct {
		Username  string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email     string `form:"email" json:"email" binding:"exists,email"`
		Password  string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		FirstName string `form:"first_name" json:"first_name"`
		LastName  string `form:"last_name" json:"last_name"`
		Phone     string `form:"phone" json:"phone"`
	} `json:"user"`
	UserModel models.User `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)

	if err != nil {
		log.WithFields(log.Fields{
			"self": self,
		}).Info("Error parsing", err)

		return err
	}
	self.UserModel.UserName = self.User.Username
	self.UserModel.Email = self.User.Email
	self.UserModel.FirstName = self.User.FirstName
	self.UserModel.LastName = self.User.LastName
	self.UserModel.Phone = self.User.Phone

	if self.User.Password != common.NBRandomPassword {
		self.UserModel.SetPassword(self.User.Password)
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
	userModelValidator.User.Username = userModel.UserName
	userModelValidator.User.Email = userModel.Email
	userModelValidator.User.FirstName = userModel.FirstName
	userModelValidator.User.LastName = userModel.LastName
	userModelValidator.User.Phone = userModel.Phone

	return userModelValidator
}
