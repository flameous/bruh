package problems

import (
// "fmt"
)

// https://leetcode.com/problems/majority-element/
// problem desc:
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// friend of mine said that this problem could be solved in linear time
// 'canonical solution' - https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm

// my approach:
// 1. majority element appears >= N/2 times
// 2. that means if we sorted an array, we should pick the N/2'th element and it would be majority element
// 3. but sorting an array takes O(log_n * N) time at least
// 4. but we can find N'th largest element in linear time with 'quick-select' algorithm
// so final idea:
// find N'th element by quick-select algorithm and it would be the majority element
// *applause*

// time: O(n), extra space: O(1)
func majorityElement(elems []int) int {
	res := quickselect(elems, len(elems)/2)
	return res
}

func quickselect(elems []int, order int) int {
	// 'sort' in decreasing order

	low := 0
	high := len(elems) - 1

	for low <= high {

		//		fmt.Printf("1. starting partitioning: %v [%d %d]\n", elems, low, high)
		res := partition(elems, low, high)
		//		fmt.Printf("partitioned result: %v, pos_idx: %d. looking for idx = %d\n", elems, res, order)

		if res == order {
			return elems[res]
		}

		if res > order {
			high = res - 1
		} else {
			// res < order
			low = res + 1
		}
	}

	return -1 // maybe panic?
}

// partition in-place
func partition(elems []int, low, high int) int {
	i := low - 1

	for j := low; j < high; j++ {
		if elems[j] >= elems[high] {
			i++
			elems[i], elems[j] = elems[j], elems[i]
		}
	}

	elems[i+1], elems[high] = elems[high], elems[i+1]
	return i + 1
}
