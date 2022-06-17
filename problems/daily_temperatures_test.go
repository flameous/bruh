package problems

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dailyTemperatures(t *testing.T) {
	var cases = []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{1, 2, 3, 4, 5},
			out: []int{1, 1, 1, 1, 0},
		},
		{
			in:  []int{9},
			out: []int{0},
		},
		{
			in:  []int{13, 15, 9, 8, 12, 17},
			out: []int{1, 4, 2, 1, 1, 0},
		},
	}

	for _, tc := range cases {
		res := dailyTemperatures(tc.in)

		require.True(t, reflect.DeepEqual(tc.out, res))
	}
}
