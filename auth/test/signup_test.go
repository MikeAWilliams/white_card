package testauth

import (
	"errors"
	"testing"

	"github.com/MikeAWilliams/white_card/auth"
	"github.com/stretchr/testify/require"
)

func TestPasswordInvalid(t *testing.T) {
	db := getTestDB([]auth.User{
		auth.User{Email: "e@example.com", Password: "password"}},
		nil, nil, nil)

	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	err := auth.Signup(newUser, &db, func(_ string) error { return errors.New("bad bad bad") })

	require.NotNil(t, err)
	require.Contains(t, err.Message, "bad bad bad")
}
