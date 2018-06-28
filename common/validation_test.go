package common

import (
	//"fmt"
	"testing"
)

// testing the validation package
func TestBind(ts *testing.T) {
	user := &UserModelValidator{Username: "",
		Email:     "test@gmail",
		Password:  "testing",
		FirstName: "",
		LastName:  "",
		Phone:     "",
	}
	err := Bind(user)
	//fmt.Println(err)
	if err == nil {
		ts.Errorf("validation for blank User name failed")
	}
}
