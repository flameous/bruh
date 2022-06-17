package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecompressString(t *testing.T) {
	t.Run("regular string", func(t *testing.T) {
		expected := "ABCDE"
		got := decompress("ABCDE")

		require.Equal(t, expected, got)
	})

	t.Run("plain compressed string", func(t *testing.T) {
		expected := "ABBBCCD"
		got := decompress("A3B2CD")

		require.Equal(t, expected, got)
	})

	t.Run("nested string", func(t *testing.T) {
		expected := "ABBBCDDDECDDDEF"
		got := decompress("A3B2(C3DE)F")

		require.Equal(t, expected, got)
	})

	t.Run("nested string #2", func(t *testing.T) {
		expected := "ABBCDEFFFDEFFFGCDEFFFDEFFFGHII"
		got := decompress("A2B2(C2(DE3F)G)H2I")

		require.Equal(t, expected, got)
	})

	t.Run("nested string", func(t *testing.T) {
		expected := "ABCDABCDABCDABCD"
		got := decompress("4(ABCD)")

		require.Equal(t, expected, got)
	})

}
