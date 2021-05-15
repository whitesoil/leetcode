package leetcode

func ThreeSum(nums []int) [][]int {
	var output [][]int

	if len(nums) < 3 {
		return output
	}

	sortedNums := quicksort(nums)
	length := len(sortedNums)

	for i := 0; i < length-2; {
		start := i + 1
		end := length - 1

		for start < end {
			sum := sortedNums[i] + sortedNums[start] + sortedNums[end]

			switch {
			case sum == 0:
				output = append(output, []int{sortedNums[i], sortedNums[start], sortedNums[end]})
				end--

				x := start
				for sortedNums[x] == sortedNums[start] && x < length-2 {
					x++
				}

				if x == start {
					start++
				} else {
					start = x
				}
			case sum > 0:
				end--
			case sum < 0:
				start++
			}
		}

		j := i
		for sortedNums[j] == sortedNums[i] && j < length-2 {
			j++
		}

		if i == j {
			i++
		} else {
			i = j
		}
	}

	return output
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
