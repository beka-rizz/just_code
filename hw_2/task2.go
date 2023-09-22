package hw_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	cur1, cur2 := list1, list2

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if cur1.Val <= cur2.Val {
		head = cur1
		cur1 = cur1.Next
	} else {
		head = cur2
		cur2 = cur2.Next
	}

	cur := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	if cur1 == nil {
		cur.Next = cur2
	} else {
		cur.Next = cur1
	}
	return head
}
