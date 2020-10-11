package auth

import (
	e "github.com/MikeAWilliams/white_card/wraperr"
)

type PasswordValidator func(s string) error

func AddUser(u User, db Database, pwIsValid PasswordValidator) *e.WrapError {
	pwErr := pwIsValid(u.Password)
	if nil != pwErr {
		return e.Wrap(pwErr, "The password is invalid %v", pwErr.Error())
	}

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
