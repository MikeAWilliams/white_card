package testauth

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/white_card/auth"
)

type testDB struct {
}

func (db *testDB) GetUser(email string) (auth.User, error) {
	return auth.User{}, nil
}

func (db *testDB) AddUser(user auth.User) error {
	return nil
}

func (db *testDB) UserExists(email string) (bool, error) {
	return false, nil
}

func TestAddUser(t *testing.T) {
	db := testDB{}
	err := auth.AddUser("e@example.com", &db)

	require.NotNil(t, err)
}
