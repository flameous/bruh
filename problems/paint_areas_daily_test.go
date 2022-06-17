package problems

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaintAmountEachDay(t *testing.T) {
	in := []area{{start: 1, end: 3}, {start: 5, end: 8}, {start: 2, end: 6}, {start: 1, end: 8}}

	result := paintAmountEachDay(in)

	expected := []int{2, 3, 2, 0}
	require.True(t, reflect.DeepEqual(result, expected))
}
