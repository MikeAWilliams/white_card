package testauth

import "github.com/MikeAWilliams/white_card/auth"

type testDB struct {
	table           []auth.User
	userExistsError error
	addUserError    error
	getUserError    error
}

func (db *testDB) GetUser(email string) (auth.User, error) {
	return auth.User{}, nil
}

func (db *testDB) AddUser(user auth.User) error {
	if nil != db.addUserError {
		return db.addUserError
	}
	db.table = append(db.table, user)
	return nil
}

func (db *testDB) UserExists(email string) (bool, error) {
	if nil != db.userExistsError {
		return false, db.userExistsError
	}
	for _, usr := range db.table {
		if usr.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func getTestDB(users []auth.User, existErr error, addError error, getError error) testDB {
	result := testDB{}
	result.table = users
	result.userExistsError = existErr
	result.addUserError = addError
	result.getUserError = getError
	return result
}
