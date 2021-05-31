package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	return mergeTwoLists(MergeKLists(lists[0:length/2]), MergeKLists(lists[length/2:]))
}

func mergeTwoLists(listA *ListNode, listB *ListNode) *ListNode {
	start := &ListNode{}
	container := start

	for listA != nil || listB != nil {
		if listA == nil && listB != nil {
			container.Next = listB
			listB = listB.Next
			container = container.Next
			continue
		}

		if listA != nil && listB == nil {
			container.Next = listA
			listA = listA.Next
			container = container.Next
			continue
		}

		if listA != nil && listB != nil {
			if listA.Val < listB.Val {
				container.Next = listA
				listA = listA.Next
			} else {
				container.Next = listB
				listB = listB.Next
			}
			container = container.Next
			continue
		}
	}

	return start.Next
}
