package testauth

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/white_card/auth"
)

type testDB struct {
	table []auth.User
}

func (db *testDB) GetUser(email string) (auth.User, error) {
	return auth.User{}, nil
}

func (db *testDB) AddUser(user auth.User) error {
	return nil
}

func (db *testDB) UserExists(email string) (bool, error) {
	for _, usr := range db.table {
		if usr.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func getTestDB(users []auth.User) testDB {
	result := testDB{}
	result.table = users
	return result
}

func TestAddUser(t *testing.T) {
	db := getTestDB([]auth.User{
		auth.User{Email: "e@example.com", Password: "password"}})

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.NotNil(t, err)
}
