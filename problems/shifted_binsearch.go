package problems

// find element's index in sorted then shifted array
// return -1 if not found
func shiftedBinanySearch(arr []int, elem int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		m := (l + r) / 2

		mid, left, right := arr[m], arr[l], arr[r]
		if arr[m] == elem {
			return m
		}

		// pivot point on the 'right' side
		// 3 4 5 6 7 8 0 1 2
		// l       m  ^    r
		if left <= mid {
			if elem < mid && left <= elem {
				r = m - 1
			} else {
				l = m + 1
			}

		} else if mid <= right {
			if mid < elem && elem <= right {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
