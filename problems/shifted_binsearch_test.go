package problems

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestShiftedBinanySearch(t *testing.T) {
	rand.Seed(time.Now().Unix())

	tmp := make([]int, 100)

	prev := 0
	for i := 0; i < 100; i++ {
		num := prev + rand.Intn(10) + 1
		prev = num
		tmp[i] = num
	}

	point := rand.Intn(100) // [0, 99]

	var slice []int
	slice = append(slice, tmp[point:]...)
	slice = append(slice, tmp[:point]...)

	t.Run("find", func(t *testing.T) {
		randIdx := rand.Intn(len(slice))
		elem := slice[randIdx]

		idx := shiftedBinanySearch(slice, elem)

		t.Logf("random index = %d, elem = %d, array = %v\n", randIdx, elem, slice)

		require.Equalf(t, randIdx, idx, "%v failed", randIdx)
	})

	t.Run("find, corner case #1", func(t *testing.T) {
		exceptedIdx := 0
		elem := slice[exceptedIdx]

		idx := shiftedBinanySearch(slice, elem)

		require.Equal(t, exceptedIdx, idx)
	})

	t.Run("find, corner case #2", func(t *testing.T) {
		exceptedIdx := len(slice) - 1
		elem := slice[exceptedIdx]

		idx := shiftedBinanySearch(slice, elem)

		require.Equal(t, exceptedIdx, idx)
	})

	t.Run("find, corner case #3", func(t *testing.T) {
		exceptedIdx := rand.Intn(len(tmp))
		elem := tmp[exceptedIdx]

		idx := shiftedBinanySearch(tmp, elem)

		require.Equal(t, exceptedIdx, idx)
	})

	t.Run("not found", func(t *testing.T) {
		elem := tmp[0] - 1

		idx := shiftedBinanySearch(slice, elem)

		require.Equal(t, -1, idx)
	})
}
