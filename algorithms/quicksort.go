package algorithms

// (avg) time complexity = O(n * log_2 n)
// space complexity = O(n)
// TODO: quicksort in-place
func quicksort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	pivotIdx := len(data) / 2
	var less, greater []int
	for i := 0; i < len(data); i++ {
		if i != pivotIdx {
			if v := data[i]; v < data[pivotIdx] {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
	}

	ret := make([]int, 0, len(data))

	ret = append(ret, quicksort(less)...)
	ret = append(ret, data[pivotIdx])
	ret = append(ret, quicksort(greater)...)

	return ret
}
