package problems

import (
	"testing"
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

		if len(res) != len(tc.out) {
			t.Fatalf("expected len = %d, but got = %d", len(tc.out), len(res))
		}

		for i := 0; i < len(res); i++ {
			if res[i] != tc.out[i] {
				t.Fatalf("expected val = %d, but got = %d", res[i], tc.out[i])
			}
		}
	}
}
