package auth

import "fmt"

func AddUser(u User, db Database) error {
	exists, err := db.UserExists(u.Email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("A user with email %v already exists", u.Email)
	}
	return nil
}
