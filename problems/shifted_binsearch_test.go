package problems

import (
	"math/rand"
	"testing"
	"time"
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

		if idx == -1 {
			t.Fatalf("expected idx = %d, but got -1 (%d %d)", randIdx, elem, slice[randIdx])
		}

		if randIdx != idx {
			t.Fatalf("expected %d (value = %d), but got %d idx (value = %d)", randIdx, elem, idx, slice[idx])
		}
	})

	t.Run("find, corner case #1", func(t *testing.T) {
		exceptedIdx := 0
		elem := slice[exceptedIdx]

		idx := shiftedBinanySearch(slice, elem)

		if idx == -1 {
			t.Fatalf("expected idx = %d, but got -1 (%d %d)", exceptedIdx, elem, slice[exceptedIdx])
		}

		if exceptedIdx != idx {
			t.Fatalf("expected %d (value = %d), but got %d idx (value = %d)", exceptedIdx, elem, idx, slice[idx])
		}
	})

	t.Run("find, corner case #2", func(t *testing.T) {
		exceptedIdx := len(slice) - 1
		elem := slice[exceptedIdx]

		idx := shiftedBinanySearch(slice, elem)

		if idx == -1 {
			t.Fatalf("expected idx = %d, but got -1 (%d %d)", exceptedIdx, elem, slice[exceptedIdx])
		}

		if exceptedIdx != idx {
			t.Fatalf("expected %d (value = %d), but got %d idx (value = %d)", exceptedIdx, elem, idx, slice[idx])
		}
	})

	t.Run("find, corner case #3", func(t *testing.T) {
		exceptedIdx := rand.Intn(len(tmp))
		elem := tmp[exceptedIdx]

		idx := shiftedBinanySearch(tmp, elem)

		if idx == -1 {
			t.Fatalf("expected idx = %d, but got -1 (%d %d)", exceptedIdx, elem, tmp[exceptedIdx])
		}

		if exceptedIdx != idx {
			t.Fatalf("expected %d (value = %d), but got %d idx (value = %d)", exceptedIdx, elem, idx, tmp[idx])
		}
	})

	t.Run("not found", func(t *testing.T) {
		elem := tmp[0] - 1

		idx := shiftedBinanySearch(slice, elem)

		if idx != -1 {
			t.Fatalf("expected -1, but got %d idx (value = %d)", idx, slice[idx])
		}
	})
}
