package emit

import (
	"os"
	"strings"
	"testing"

	"github.com/codeliger/tygo/tygo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmit(t *testing.T) {
	t.Parallel()

	goCodeBytes, err := os.ReadFile("emit.txt")
	require.NoError(t, err)

	goCode := string(goCodeBytes)

	pkgConfig := &tygo.PackageConfig{
		Path: "emit",
	}

	tsCode, err := tygo.GenerateFromString(goCode, pkgConfig)
	require.NoError(t, err)

	expectedBytes, err := os.ReadFile("index.ts")
	require.NoError(t, err)

	expected := string(expectedBytes)

	// Normalize line endings
	tsCode = strings.ReplaceAll(tsCode, "\r\n", "\n")
	expected = strings.ReplaceAll(expected, "\r\n", "\n")

	assert.Equal(t, expected, tsCode)
}
