package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabase(t *testing.T) {
	_, err := InitDB()
	require.NoError(t, err)

}
