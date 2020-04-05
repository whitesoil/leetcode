package leetcode

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers is 2nd problem
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2

	node := new(ListNode)
	result := node

	for {
		val1 := 0
		val2 := 0

		if n1 != nil {
			val1 = n1.Val
			n1 = n1.Next
		}

		if n2 != nil {
			val2 = n2.Val
			n2 = n2.Next
		}

		sum := val1 + val2 + node.Val

		if n1 == nil && n2 == nil {
			if sum >= 10 {
				node.Next = new(ListNode)
				node.Val = sum - 10
				node.Next.Val = 1
			} else {
				node.Val = sum
			}

			break
		} else {
			node.Next = new(ListNode)

			if sum >= 10 {
				node.Val = sum - 10
				node.Next.Val = 1
			} else {
				node.Val = sum
			}

			node = node.Next
		}
	}

	return result
}
