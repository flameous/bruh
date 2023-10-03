package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseWords(t *testing.T) {
	data := []rune("hello world in go")

	reverseWords(data)

	require.Equal(t, "go in world hello", string(data))
}
