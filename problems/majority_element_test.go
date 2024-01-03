package problems

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	// corner cases:
	// one element
	result := majorityElement([]int{1})
	require.Equal(t, 1, result)

	// two elements
	result = majorityElement([]int{1, 1})
	require.Equal(t, 1, result)

	// three elements
	result = majorityElement([]int{1, 1, 2})
	require.Equal(t, 1, result)

	// four elements
	result = majorityElement([]int{1, 1, 2, 3})
	require.Equal(t, 1, result)

	elems := []int{3, 8, 3, 3, 2, 3, 1, 6}
	result = majorityElement(elems)
	require.Equal(t, 3, result)

	// random
	randInt := rand.Intn(1000000)
	sl := createSliceWithMajorityElement(randInt)
	res := majorityElement(sl)
	require.Equal(t, randInt, res)
}

// helper
func createSliceWithMajorityElement(majorityElement int) []int {
	size := 1 + rand.Intn(100)

	ret := make([]int, 0, size)

	for i := 0; i < size; i++ {
		if i < size/2+1 {
			ret = append(ret, majorityElement)
		} else {
			ret = append(ret, majorityElement+i)
		}
	}

	rand.Shuffle(len(ret), func(i, j int) { ret[i], ret[j] = ret[j], ret[i] })

	return ret
}
