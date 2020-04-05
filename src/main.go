package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	medianIndex := float64(totalLength) / 2.0
	var ceilMedianIndex float64
	var floorMedianIndex float64

	if (totalLength % 2) == 0 {
		ceilMedianIndex = math.Ceil(medianIndex)
		floorMedianIndex = ceilMedianIndex - 1.0
	} else {
		ceilMedianIndex = math.Floor(medianIndex)
		floorMedianIndex = ceilMedianIndex
	}

	leftValue := 0
	rightValue := 0

	for ceilMedianIndex >= 0 || floorMedianIndex >= 0 {
		if len(nums1) == 0 {
			if floorMedianIndex == 0 {
				leftValue = nums2[0]
			}

			if ceilMedianIndex == 0 {
				rightValue = nums2[0]
			}

			if len(nums2) > 1 {
				nums2 = nums2[1:]
			} else {
				nums2 = []int{}
			}
		} else if len(nums2) == 0 {
			if floorMedianIndex == 0 {
				leftValue = nums1[0]
			}

			if ceilMedianIndex == 0 {
				rightValue = nums1[0]
			}

			if len(nums1) > 1 {
				nums1 = nums1[1:]
			} else {
				nums1 = []int{}
			}
		} else if nums1[0] >= nums2[0] {
			if floorMedianIndex == 0 {
				leftValue = nums2[0]
			}

			if ceilMedianIndex == 0 {
				rightValue = nums2[0]
			}

			if len(nums2) > 1 {
				nums2 = nums2[1:]
			} else {
				nums2 = []int{}
			}
		} else if nums2[0] > nums1[0] {
			if floorMedianIndex == 0 {
				leftValue = nums1[0]
			}

			if ceilMedianIndex == 0 {
				rightValue = nums1[0]
			}

			if len(nums1) > 1 {
				nums1 = nums1[1:]
			} else {
				nums1 = []int{}
			}
		}

		floorMedianIndex--
		ceilMedianIndex--
	}

	return float64(leftValue+rightValue) / 2.0
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	print(findMedianSortedArrays(nums1, nums2))
}
