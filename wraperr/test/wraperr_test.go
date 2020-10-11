package testwraperr

import (
	"errors"
	"testing"

	e "github.com/MikeAWilliams/white_card/wraperr"
	"github.com/stretchr/testify/require"
)

func TestAddUserExists(t *testing.T) {
	testObject := e.Wrap(errors.New("somethign bad happened"), "V1=%v, V2=%v", "A bad thing", 2)
	require.Equal(t, "V1=A bad thing, V2=2", testObject.Message)
}
