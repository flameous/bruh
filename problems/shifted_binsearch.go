package problems

// find element's index in sorted then shifted array
// return -1 if not found
func shiftedBinanySearch(arr []int, elem int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == elem {
			return mid
		}

		if arr[mid] > elem {
			if arr[l] <= elem {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if arr[mid] < elem {
			if arr[r] >= elem {
				l = mid + 1
			} else {
				r = mid - 1 // not covered, bug is probably here
			}
		}
	}

	return -1
}
