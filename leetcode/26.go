package leetcode

import "sort"

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	const Inf = 10001

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == Inf {
			continue
		}

		if nums[i] == nums[i+1] {
			nums[i] = Inf
		}
	}

	sort.Ints(nums)

	length := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == Inf {
			break
		}

		length++
	}

	return length
}
