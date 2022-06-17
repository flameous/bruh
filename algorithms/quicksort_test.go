package algorithms

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuicksort(t *testing.T) {
	d := rand.Perm(100)
	quicksort(d)

	require.IsIncreasing(t, d)
}
