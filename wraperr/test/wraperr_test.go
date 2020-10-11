package testwraperr

import (
	"errors"
	"testing"

	e "github.com/MikeAWilliams/white_card/wraperr"
	"github.com/stretchr/testify/require"
)

func TestErrorWithMessageArgs(t *testing.T) {
	testObject := e.Wrap(errors.New("something bad happened"), "V1=%v, V2=%v", "A bad thing", 2)

	require.Equal(t, "V1=A bad thing, V2=2", testObject.Message)
	require.Equal(t, "something bad happened", testObject.Inner.Error())
}
func TestErrorWithSimpleMessage(t *testing.T) {
	testObject := e.Wrap(errors.New("something bad happened"), "Here is a simple message")

	require.Equal(t, "Here is a simple message", testObject.Message)
}

func TestMakeError(t *testing.T) {
	testObject := e.Make("something bad happened for the %vnd time", 2)

	require.Equal(t, "something bad happened for the 2nd time", testObject.Message)
	require.False(t, testObject.HasInnerError())
}

func TestItIsAnError(t *testing.T) {
	var testObject interface{} = e.Make("message")
	err, ok := testObject.(error)

	require.True(t, ok)
	require.Equal(t, "message", err.Error())
}
