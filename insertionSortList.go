package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToString() []int {
	var a []int
	tmp := l
	for tmp != nil {
		a = append(a, tmp.Val)
		tmp = tmp.Next
	}
	return a
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	sorted := &ListNode{}
	{
		cur := head
		for ok := true; ok; ok = cur != nil {
			insertionSortListSort(sorted, cur)
			cur = cur.Next

		}
	}
	return sorted.Next
}

func insertionSortListSort(sorted, cur *ListNode) {
	prev := sorted
	for ok := true; ok; ok = prev.Next != nil {
		if prev.Next == nil {
			prev.Next = &ListNode{
				Val:  cur.Val,
				Next: nil,
			}
			return
		}
		if prev.Next.Val > cur.Val {
			prev.Next = &ListNode{
				Val:  cur.Val,
				Next: prev.Next,
			}
			return
		}
		prev = prev.Next
	}
	prev.Next = &ListNode{
		Val:  cur.Val,
		Next: nil,
	}
}
