package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// Helper Functions
func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

// Define interface
type User interface {
	Authorized() int
	ValidatePassword(string) error
}

type Admin struct {
	Name  string
	Email string
}

type AdminWithPassword struct {
	*Admin
	Password string
}

func (this Admin) Authorized() int {
	return 5
}

func (this *AdminWithPassword) ValidatePassword(password string) error {
	errValidatePassword := verifyPassword(password)
	if errValidatePassword == nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			return err
		}
		this.Password = string(bytes)
		return nil
	} else {
		return errValidatePassword
	}
}

func Auth(user User, password string) bool {
	if user.Authorized() > 4 && user.ValidatePassword(password) == nil {
		return true
	}
	return false
}

func main() {
	admin := Admin{"Admin", "admin@mail.com"}
	adminWpassdors := AdminWithPassword{&admin, ""}
	password := "Apple"

	fmt.Println(Auth(&adminWpassdors, password))
}
