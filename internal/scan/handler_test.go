package scan_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mickamy/gon/internal/scan"
)

func TestScanHandler(t *testing.T) {
	t.Parallel()

	// arrange
	dir := "./testdata/scan_test_handler.go"

	// act
	handlers, err := scan.Handlers(dir)

	// assert
	require.NoError(t, err)
	expected := []string{
		"ListUser",
		"GetUser",
		"CreateUser",
		"UpdateUser",
		"DeleteUser",
	}
	require.Equal(t, expected, handlers)
}
