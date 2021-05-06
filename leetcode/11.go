package leetcode

import "math"

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104
func MaxArea(height []int) int {
	start := 0
	end := len(height) - 1
	area := 0.0

	for start < end {
		y := math.Min(float64(height[start]), float64(height[end]))

		area = math.Max(y*float64(end-start), area)

		for height[start] <= int(y) && start < end {
			start++
		}

		for height[end] <= int(y) && start < end {
			end--
		}
	}

	return int(area)
}
