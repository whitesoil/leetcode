package leetcode

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	combinations := kSum(nums, target, 4)
	return combinations
}

func kSum(nums []int, target int, k int) [][]int {
	output := [][]int{}

	if len(nums) == 0 || nums[0]*k > target || nums[len(nums)-1]*k < target {
		return output
	}

	if k == 2 {
		lo := 0
		hi := len(nums) - 1

		for lo < hi {
			sum := nums[lo] + nums[hi]

			if sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]) {
				hi--
			} else if sum < target || (lo > 0 && nums[lo] == nums[lo-1]) {
				lo++
			} else {
				newCombi := []int{nums[lo], nums[hi]}
				output = append(output, newCombi)
				lo++
				hi--
			}
		}

		return output
	}

	for i, v := range nums {
		if i == 0 || nums[i-1] != nums[i] {
			combinations := kSum(nums[i+1:], target-v, k-1)

			for _, combi := range combinations {
				newCombi := append([]int{v}, combi...)
				output = append(output, newCombi)
			}
		}
	}

	return output
}
