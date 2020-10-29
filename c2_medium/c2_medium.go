package c2_medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	curSum := l1.Val + l2.Val
	next := addTwoNumbers(l1.Next, l2.Next)

	if curSum >= 10 {
		curSum %= 10
		next = addTwoNumbers(next, &ListNode{
			Val: 1,
		})
	}

	return &ListNode{
		Val:  curSum,
		Next: next,
	}
}
