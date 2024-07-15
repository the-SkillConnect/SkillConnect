package graphql

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// Constants for validation
const (
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

// CreateUserParams struct
type CreateUserParams struct {
	FirstName     string
	LastName      string
	Password      string
	Email         string
	MobilePhone   string
	WalletAddress string
}

// Validate function for CreateUserParams
func (params CreateUserParams) Validate() map[string]string {
	errors := make(map[string]string, 5)

	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("first name length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("last name length should be at least %d characters", minLastNameLen)
	}
	if !isPasswordValid(params.Password) {
		errors["password"] = "password must be at least 8 characters long and include both letters and numbers"
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email %s is invalid", params.Email)
	}
	if len(params.MobilePhone) == 0 {
		errors["mobilePhone"] = "mobile phone is required"
	}
	if len(params.WalletAddress) == 0 {
		errors["walletAddress"] = "wallet address is required"
	}

	return errors
}

func isPasswordValid(p string) bool {
	passwordRegex := regexp.MustCompile(`^[A-Za-z\d]{8,}$`)
	return passwordRegex.MatchString(p)
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

// IsPasswordValid function to compare hashed password with plain password
func IsPasswordValid(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}
