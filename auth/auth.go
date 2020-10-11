package auth

import (
	e "github.com/MikeAWilliams/white_card/wraperr"
)

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
