package leetcode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := 10000

	for i, v := range nums {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := v + nums[left] + nums[right]

			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}

			if diff == 0 {
				return target - diff
			}

			if sum < target {
				left++
			} else {
				right--
			}

		}
	}

	return target - diff
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// func quicksort2(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}

// 	end := len(arr) - 1
// 	pivot := len(arr) / 2

// 	arr[end], arr[pivot] = arr[pivot], arr[end]

// 	left := 0
// 	right := len(arr) - 2

// 	for left <= right {
// 		if arr[left] > arr[end] {
// 			arr[left], arr[right] = arr[right], arr[left]
// 			right--
// 		} else {
// 			left++
// 		}
// 	}

// 	arr[end], arr[left] = arr[left], arr[end]

// 	quicksort2(arr[0:pivot])
// 	quicksort2(arr[pivot:end])

// 	return arr
// }
