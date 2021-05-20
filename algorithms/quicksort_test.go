package algorithms

import (
	"math/rand"
	"testing"
)

func TestQuicksort(t *testing.T) {
	d := rand.Perm(100)

	sorted := quicksort(d)
	t.Log(sorted)
	for i := 0; i < len(sorted); i++ {
		if i != sorted[i] {
			t.Fatalf("expected value %d but got %d", i, sorted[i])
		}
	}
}
