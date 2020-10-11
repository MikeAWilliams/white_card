package testauth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/white_card/auth"
)

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
	exists, _ := db.UserExists("e@example.com")
	require.True(t, exists)
}

func TestAddUserErrorOnAdd(t *testing.T) {
	db := getTestDB([]auth.User{}, nil, errors.New("an add error"), nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.AddUser(newUser, &db)

	require.NotNil(t, err)
}
