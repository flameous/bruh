package algorithms

import (
	"math/rand"
	"testing"
)

func TestQuicksort(t *testing.T) {
	d := rand.Perm(100)
	quicksort(d)

	t.Log(d)
	for i := 0; i < len(d); i++ {
		if i != d[i] {
			t.Fatalf("expected value %d but got %d", i, d[i])
		}
	}
}
