package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{10, 15, 5},
			out: 15,
		},
		{
			in:  []int{1, 5, 2, 6, 8, 23, 5, 6, 8, 5, 673, 6, 4, 6, 7, 12, 6, 1, 9, 62, 1},
			out: 60,
		},
	}

	for _, tc := range testCases {
		res := minCostClimbingStairs(tc.in)

		require.Equal(t, tc.out, res)
	}
}
