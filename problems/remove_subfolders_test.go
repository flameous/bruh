package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveSubfolders(t *testing.T) {
	in := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	expected := []string{"/a", "/c/d", "/c/f"}

	ret := removeSubfolders(in)

	require.ElementsMatch(t, expected, ret)
}
