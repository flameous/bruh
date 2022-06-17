package problems

// time complexity: O(n)
// extra space: O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {

		for {
			if len(stack) == 0 {
				break
			}

			idx := stack[len(stack)-1]
			if temperatures[idx] < temperatures[i] {
				stack = stack[:len(stack)-1]
				res[idx] = i - idx
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	return res
}
