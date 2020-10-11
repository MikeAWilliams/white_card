package auth

import (
	"fmt"

	e "github.com/MikeAWilliams/white_card/wraperr"
)

type PasswordValidator func(pw string) error

type EmailSender interface {
	Send(address string, subject string, body string) error
}

func AddUser(u User, db Database) *e.WrapError {
	exists, err := db.UserExists(u.Email)
	if err != nil {
		fmt.Println("error in user exists")
		return e.Wrap(err, "Database error during UserExists")
	}

	if exists {
		fmt.Println("the user already exists")
		return e.Make("A user with email %v already exists", u.Email)
	}

	addErr := db.AddUser(u)
	if addErr != nil {
		fmt.Println("adding the user made an error")
		return e.Wrap(addErr, "Database error during AddUser")
	}
	return nil
}

func Signup(u User, db Database, pwValidator PasswordValidator, email EmailSender) *e.WrapError {
	err := pwValidator(u.Password)
	if err != nil {
		return e.Wrap(err, "There is an error with the password %v", err.Error())
	}

	addUserError := AddUser(u, db)

	if addUserError != nil {
		fmt.Println("this is the shit")
		fmt.Println(addUserError.Error())
		return e.Wrap(err, "There was a problem adding the user %v", addUserError.Error())
	}

	err = email.Send(u.Email, "Validate Your Eamil Address", getValidateEmailBody(u.Email))
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
