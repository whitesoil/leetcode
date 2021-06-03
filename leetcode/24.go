package leetcode

/** 39분 시작
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	startNode := head
	curNode := head
	prevNode := head
	for i := 0; curNode.Next != nil; i++ {
		if i%2 == 0 {
			next := curNode.Next
			curNode.Next = next.Next
			next.Next = curNode

			if i == 0 {
				startNode = next
				prevNode = next
			} else {
				prevNode.Next = next
			}
		} else {
			prevNode = curNode
			curNode = curNode.Next
		}
	}

	return startNode
}
