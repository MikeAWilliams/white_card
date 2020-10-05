package testauth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/white_card/auth"
)

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

func TestAddUserExists(t *testing.T) {
	db := getTestDB([]auth.User{
		auth.User{Email: "e@example.com", Password: "password"}},
		nil, nil, nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.NotNil(t, err)
}

func TestAddUserDBError(t *testing.T) {
	db := getTestDB([]auth.User{}, errors.New("some user error"), nil, nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.NotNil(t, err)
}

func TestAddUserUserIsAdded(t *testing.T) {
	db := getTestDB([]auth.User{}, nil, nil, nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.Nil(t, err)
	exists, err := db.UserExists("e@example.com")
	require.True(t, exists)
}

func TestAddUserErrorOnAdd(t *testing.T) {
	db := getTestDB([]auth.User{}, nil, errors.New("an add error"), nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.NotNil(t, err)
}
