package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h, _ := revertHeadTail(head)
	return h
}

func revertHeadTail(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	h, t := revertHeadTail(head.Next)
	t.Next = head
	head.Next = nil

	return h, head
}
