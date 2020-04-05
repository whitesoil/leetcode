package leetcode

// TwoSum is 1st problem
func TwoSum(nums []int, target int) []int {
	ints := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		ints[nums[i]] = i

		if val, exist := ints[(target - nums[i+1])]; exist {
			return []int{val, i + 1}
		}
	}

	return []int{}
}
