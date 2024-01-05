package problems

// 1004. Max Consecutive Ones III
// desc:
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

import (
//    "fmt"
)

func longestOnes(nums []int, k int) int {
	var endIdx, curLen, maxLen int

    startIdx := -1

	remained := k

	for endIdx < len(nums) {
        v := nums[endIdx]
        
		// algorithm:
        // move right pointer until we reached out the limit of remained flips
		if v == 1 || remained > 0 {
            // fmt.Printf("%v [%d - %d], curr value = '%d', remained = %d\n", nums[:endIdx+1], startIdx, endIdx+1, v, remained)
			if v == 0 {
				remained--
                // fmt.Println("remained--")
			}

			curLen = endIdx - startIdx
			if curLen > maxLen {
				maxLen = curLen
			}
            endIdx++
			continue
		}

		// v == 0 and remained == 0
        // move left pointer until we restore at least one remained flip 
		for remained == 0 && startIdx <= endIdx {
            // fmt.Printf("else: cur idx = %d (startIdx = %d), value = %d\n", endIdx, startIdx, v)
			startIdx++
			if nums[startIdx] == 0 {
				remained++
			}

			if remained == k {
				break
			}
		}
	}

	return maxLen
}
