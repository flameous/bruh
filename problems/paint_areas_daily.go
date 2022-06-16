package problems

// input:
// 1D area (line)
// array of jobs (needed to be done by order)
// job is described with two dots - start and end of area needed to paint
// calculate amount of paint needed to process each job
// !! if some area already painted, no need to paint it again !!

// example: [[1, 5], [3, 6]] -> [4, 1]
// first day area - [1, 5] -> 4
// second day area - [3, 6] -> area from 3 to 5 already painted -> need to paint only [5, 6] area -> 1

type area struct {
	start, end int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func overlap(a1, a2 area) int {
	maxStart := max(a1.start, a2.start)
	minEnd := min(a1.end, a2.end)
	if maxStart > minEnd {
		return 0
	}
	return minEnd - maxStart
}

func merge(a1, a2 area) area {
	return area{
		start: min(a1.start, a2.start),
		end:   max(a1.end, a2.end),
	}
}

// time complexity: O(n**2)
// extra space complexity: O(n)
func paintAmountEachDay(jobs []area) []int {
	// todo: reuse 'jobs' array to store overlaped areas
	paintedAreas := []area{}

	result := make([]int, 0, len(jobs))

	for _, job := range jobs {
		paint := job.end - job.start

		for idx := 0; idx < len(paintedAreas); {
			ovl := overlap(job, paintedAreas[idx])
			if ovl > 0 {
				paint -= ovl
				job = merge(job, paintedAreas[idx])

				// pop element (swap with last element, pop last element)
				paintedAreas[idx], paintedAreas[len(paintedAreas)-1] = paintedAreas[len(paintedAreas)-1], paintedAreas[idx]
				paintedAreas = paintedAreas[:len(paintedAreas)-1]
			} else {
				idx++
			}
		}

		paintedAreas = append(paintedAreas, job)
		result = append(result, paint)
	}

	return result
}
