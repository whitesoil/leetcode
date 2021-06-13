package leetcode

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	right := len(nums) - 1
	left := 0

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}

	return left
}
