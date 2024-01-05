package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestOnes(t *testing.T) {
	res := longestOnes([]int{1, 1, 1, 0}, 1)
	require.Equal(t, 4, res)

	res = longestOnes([]int{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 1)
	require.Equal(t, 7, res)

	res = longestOnes([]int{1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1}, 3)
	require.Equal(t, 13, res)

	res = longestOnes([]int{1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}, 3)
	require.Equal(t, 15, res)

	res = longestOnes([]int{1, 1, 1, 1, 1}, 0) 
	require.Equal(t, 5, res)

	res = longestOnes([]int{1, 1, 0, 1, 1, 1, 0, 1}, 0) 
	require.Equal(t, 3, res)

}
