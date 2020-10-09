package c83_easy

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil {
        next := current.Next
        for next != nil && next.Val == current.Val {
            current.Next = next.Next
            next = current.Next
        }
        current = current.Next
    }
    return head
}
