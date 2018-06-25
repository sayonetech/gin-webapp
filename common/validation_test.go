package common

import (
	"testing"
)

// testing the validation package
func TestBind(ts *testing.T) {
	user := &validators.UserModelValidator{Username: "",
		Email:     "test@gmail",
		Password:  "testing",
		FirstName: "",
		LastName:  "",
		Phone:     "",
		UserModel: nil}
	err := Bind(user)
	fmt.Println(err)
	if err == nil {
		t.Errorf("validation for User name failed")
	}
}
