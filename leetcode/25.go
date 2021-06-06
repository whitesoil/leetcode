package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	startNode := &ListNode{0, head}
	curNode := startNode

	for i := 0; curNode.Next != nil; i++ {
		if i%k == 0 {
			reversedHead := reverse(curNode.Next, k)

			if reversedHead == nil {
				return startNode.Next
			}

			curNode.Next = reversedHead
		}

		curNode = curNode.Next
	}

	return startNode.Next
}

func reverse(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	list := make([]*ListNode, 0, k)
	curNode := head

	for i := 0; i < k; i++ {
		if curNode == nil {
			return nil
		}

		list = append(list, curNode)
		curNode = curNode.Next
	}

	nextNode := list[len(list)-1].Next

	for i := len(list) - 1; i > 0; i-- {
		list[i].Next = list[i-1]
	}

	list[0].Next = nextNode

	return list[len(list)-1]
}
