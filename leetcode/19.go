package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	count := 0
	window := make([]*ListNode, 0, n+1)

	for curNode := head; curNode != nil; curNode = curNode.Next {
		if len(window) > n {
			window = window[1:]
		}

		window = append(window, curNode)

		count++
	}

	if count == n {
		return head.Next
	}

	window[0].Next = window[1].Next

	return head
}
