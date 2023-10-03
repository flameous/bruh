package problems

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaintAmountEachDay_Case1(t *testing.T) {
	in := []area{{start: 1, end: 6}, {start: 5, end: 8}}

	result := paintAmountEachDay(in)

	expected := []int{5, 2}
	require.Truef(t, reflect.DeepEqual(result, expected), "expected: %v, got: %v", expected, result)
}

func TestPaintAmountEachDay_Case2(t *testing.T) {
	in := []area{{start: 1, end: 3}, {start: 5, end: 8}, {start: 2, end: 6}, {start: 1, end: 8}}

	result := paintAmountEachDay(in)

	expected := []int{2, 3, 2, 0}
	require.Truef(t, reflect.DeepEqual(result, expected), "expected: %v, got: %v", expected, result)
}
