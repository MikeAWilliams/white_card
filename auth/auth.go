package auth

import (
	"fmt"

	e "github.com/MikeAWilliams/white_card/wraperr"
)

type PasswordValidator func(pw string) error

type EmailSender interface {
	Send(address string, subject string, body string) error
}

type PasswordEncryptor interface {
	Encrypt(password string) (string, error)
}

type SignupDependencies struct {
	Db          Database
	PwValidator PasswordValidator
	Email       EmailSender
	Pwe         PasswordEncryptor
}

func AddUser(u User, db Database) *e.WrapError {
	exists, err := db.UserExists(u.Email)
	if err != nil {
		return e.Wrap(err, "Database error during UserExists")
	}

	if exists {
		return e.Make("A user with email %v already exists", u.Email)
	}

	addErr := db.AddUser(u)
	if addErr != nil {
		return e.Wrap(addErr, "Database error during AddUser")
	}
	return nil
}

func Signup(u User, dependencies SignupDependencies) *e.WrapError {
	err := dependencies.PwValidator(u.Password)
	if err != nil {
		return e.Wrap(err, "There is an error with the password %v", err.Error())
	}

	encryptedPw, pwEncryptError := dependencies.Pwe.Encrypt(u.Password)
	if pwEncryptError != nil {
		return e.Wrap(pwEncryptError, "There weas an error encrypting the password")
	}
	u.Password = encryptedPw
	addUserError := AddUser(u, dependencies.Db)

	if addUserError != nil {
		return e.Wrap(err, "There was a problem adding the user %v", addUserError.Error())
	}

	err = dependencies.Email.Send(u.Email, "Validate Your Eamil Address", getValidateEmailBody(u.Email))
	if nil != err {
		return e.Wrap(err, "There is an sending the validation email %v", err.Error())
	}

	return nil
}

func getValidateEmailBody(address string) string {
	link := fmt.Sprintf("%v/api/v1/auth/verify?email=%v&token=%v", getBaseURL(), address, getJWT)
	return fmt.Sprintf("Click the link below to validate your email address\r\n%v", link)
}

func getBaseURL() string {
	return "garbage//"
}

func getJWT() string {
	return "garbage"
}
