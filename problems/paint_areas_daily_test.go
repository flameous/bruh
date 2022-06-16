package problems

import "testing"

func TestPaintAmountEachDay(t *testing.T) {
	in := []area{{start: 1, end: 3}, {start: 5, end: 8}, {start: 2, end: 6}, {start: 1, end: 8}}

	result := paintAmountEachDay(in)

	if len(in) != len(result) {
		t.Fatalf("sizes dont match: %d %d", len(in), len(result))
	}

	expected := []int{2, 3, 2, 0}
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Fatalf("expected %d but got %d", expected[i], result[i])
		}
	}
}
