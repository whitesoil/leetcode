package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{}
	l3 := start

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				l3.Next = l1
				l1 = l1.Next
			} else {
				l3.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil && l2 == nil {
			l3.Next = l1
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			l3.Next = l2
			l2 = l2.Next
		}

		l3 = l3.Next
	}

	return start.Next
}
