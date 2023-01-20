package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result, head *ListNode

	switch true {
	case list1 != nil && list2 != nil:
		if list1.Val <= list2.Val {
			result = list1
			list1 = list1.Next
		} else {
			result = list2
			list2 = list2.Next
		}
	case list1 != nil:
		result = list1
		list1 = list1.Next
	case list2 != nil:
		result = list2
		list2 = list2.Next
	}
	head = result
	var r2, r3 *ListNode
	for list1 != nil || list2 != nil {
		switch true {
		case list1 != nil && list2 != nil:
			if list1.Val <= list2.Val {
				r2 = list1
				list1 = list1.Next
			} else {
				r2 = list2
				list2 = list2.Next
			}
		case list1 != nil:
			r2 = list1
			list1 = list1.Next
		case list2 != nil:
			r2 = list2
			list2 = list2.Next
		}
		r3 = result
		result.Next = r2
		result = r2
	}
	if r3 != nil && result != nil && result.Val == 0 {
		r3.Next = nil
	}
	return head
}
