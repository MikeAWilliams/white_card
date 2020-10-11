package testauth

import (
	"errors"
	"testing"

	"github.com/MikeAWilliams/white_card/auth"
	"github.com/stretchr/testify/require"
)

type emailSpy struct {
	lastAddress string
	lastSubject string
	lastBody    string
	result      error
}

func (em *emailSpy) Send(address string, subject string, body string) error {
	em.lastAddress = address
	em.lastSubject = subject
	em.lastBody = body
	return em.result
}

func TestHappyPath(t *testing.T) {
	db := getTestDB(nil, nil, nil, nil)
	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	emSpy := emailSpy{}

	err := auth.Signup(newUser, &db, func(_ string) error { return nil }, &emSpy)

	require.Nil(t, err)
	exists, _ := db.UserExists(newUser.Email)
	require.True(t, exists)
	require.Equal(t, newUser.Email, emSpy.lastAddress)
	require.Contains(t, emSpy.lastBody, "/api/v1/auth/verify?")
	require.Contains(t, emSpy.lastBody, newUser.Email)
}
func TestAddUserError(t *testing.T) {
	db := getTestDB(nil, errors.New("bad"), nil, nil)
	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	emSpy := emailSpy{}

	err := auth.Signup(newUser, &db, func(_ string) error { return nil }, &emSpy)

	require.NotNil(t, err)
}

func TestEmailError(t *testing.T) {
	db := getTestDB(nil, nil, nil, nil)
	newUser := auth.User{Email: "e@example.com", Password: "whatever"}
	emSpy := emailSpy{}
	emSpy.result = errors.New("bad email error")

	err := auth.Signup(newUser, &db, func(_ string) error { return nil }, &emSpy)

	require.NotNil(t, err)
}
func TestPasswordInvalid(t *testing.T) {
	db := getTestDB(nil, nil, nil, nil)
	newUser := auth.User{Email: "e@example.com", Password: "whatever"}

	err := auth.Signup(newUser, &db, func(_ string) error { return errors.New("bad bad bad") }, &emailSpy{})

	require.NotNil(t, err)
	require.Contains(t, err.Message, "bad bad bad")
}
