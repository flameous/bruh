package algorithms

// (avg) time complexity = O(n * log_2 n)
// space complexity = O(1)
func quicksort(data []int) {
	qsort(data, 0, len(data)-1)
}

func qsort(data []int, low, high int) {
	if len(data) < 2 {
		return
	}

	if low < high {
		p := partition(data, low, high)
		qsort(data, low, p-1)
		qsort(data, p+1, high)
	}
}

// [7 4 2 1 9 8 6] -->  7 4 2 1 9 8 | 6 (pivot)
//
// i                    i                                             i
// v                    v                                             v
//   7 4 2 1 9 8 | 6 ->   7 4 2 1 9 8 | 6 --'incr i; swap 7 and 4'--> 4 7 2 1 9 8 | 6
//   ^                      ^                                           ^
//   j                      j                                           j
//
// (then swap 7 and 2, swap 7 and 1 & so on)
//
// after all:
//     i
//     v
// 4 2 1 7 9 8 | 6
//
// 'i' is the index of the 'borderline'
// that means that every idx <= i --> (data[idx] <= pivot) and every idx > i --> (data[idx] > pivot)
// finally, swap indexes of 7 (first greater element) and pivot itself
// resulting array: [4 2 1 _6_ 9 8 7]
func partition(data []int, low, high int) int {
	i := low - 1
	for j := low; j < high; j++ {
		if data[j] <= data[high] {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[high], data[i+1] = data[i+1], data[high]
	return i + 1
}

