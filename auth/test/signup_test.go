package testauth

import (
	"errors"
	"fmt"
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

	fmt.Println(em.lastAddress)

	return em.result
}

type pwEncryptSpy struct {
	lastPassword string
	result       error
}

func (pw *pwEncryptSpy) Encrypt(password string) (string, error) {
	pw.lastPassword = password
	return password + "encrypted", pw.result
}

type testDependencies struct {
	dep auth.SignupDependencies

	db     *testDB
	emSpy  *emailSpy
	pweSpy *pwEncryptSpy
}

func getTestDependencies() testDependencies {
	result := testDependencies{}

	emSpy := emailSpy{}
	result.emSpy = &emSpy
	pwSpy := pwEncryptSpy{}
	result.pweSpy = &pwSpy

	db := getTestDB(nil, nil, nil, nil)
	result.db = &db
	result.dep.Db = &db

	result.dep.PwValidator = func(_ string) error { return nil }
	result.dep.Email = result.emSpy
	result.dep.Pwe = result.pweSpy

	return result
}

func getUser() auth.User {
	return auth.User{Email: "e@example.com", Password: "whatever"}
}

func TestHappyPath(t *testing.T) {
	dep := getTestDependencies()
	newUser := getUser()

	err := auth.Signup(newUser, dep.dep)

	require.Nil(t, err)
	exists, _ := dep.dep.Db.UserExists(newUser.Email)
	require.True(t, exists)
	require.Equal(t, newUser.Email, dep.emSpy.lastAddress)
	require.Contains(t, dep.emSpy.lastBody, "/api/v1/auth/verify?")
	require.Contains(t, dep.emSpy.lastBody, newUser.Email)

	require.Equal(t, newUser.Password, dep.pweSpy.lastPassword)
	resultUser, _ := dep.db.GetUser(newUser.Email)
	require.Contains(t, resultUser.Password, "encrypted")
}

func TestAddUserError(t *testing.T) {
	newUser := getUser()
	dep := getTestDependencies()
	dep.db.addUserError = errors.New("bad")

	err := auth.Signup(newUser, dep.dep)

	require.NotNil(t, err)
}

func TestEmailError(t *testing.T) {
	newUser := getUser()
	dep := getTestDependencies()
	dep.emSpy.result = errors.New("bad email error")

	err := auth.Signup(newUser, dep.dep)

	require.NotNil(t, err)
}

func TestPasswordInvalid(t *testing.T) {
	newUser := getUser()
	dep := getTestDependencies()
	dep.dep.PwValidator = func(_ string) error { return errors.New("bad") }

	err := auth.Signup(newUser, dep.dep)

	require.NotNil(t, err)
	require.Contains(t, err.Message, "bad")
	exists, _ := dep.db.UserExists(newUser.Email)
	require.False(t, exists)
}

func TestPasswordEncryptError(t *testing.T) {
	newUser := getUser()
	dep := getTestDependencies()
	dep.pweSpy.result = errors.New("bad")

	err := auth.Signup(newUser, dep.dep)

	require.NotNil(t, err)
	exists, _ := dep.db.UserExists(newUser.Email)
	require.False(t, exists)
}
