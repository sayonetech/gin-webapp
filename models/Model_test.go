package models

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func BenchmarkCheckPassword(bench *testing.B) {

	for i := 0; i < bench.N; i++ {

		var Utest User
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("sayone"), bcrypt.DefaultCost)
		Utest.Password = string(hashPassword)

		Utest.CheckPassword("sayone")
	}
}

func TestCheckPassword(ts *testing.T) {
	var Utest User

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("sayone"), bcrypt.DefaultCost)

	Utest.Password = string(hashPassword)

	if err := Utest.CheckPassword("sayone"); err != nil {
		ts.Error("correct passwords did not match", err)
	}

	if err := Utest.CheckPassword("sayone1"); err == nil {
		ts.Error("incorrect passwords matched", err)
	}

}
