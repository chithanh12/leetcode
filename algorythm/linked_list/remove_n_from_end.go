package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := getSize(head)
	previousIndex := size - (n + 1)

	if previousIndex < 0 {
		return head.Next
	}

	node := findNodeAt(head, previousIndex)
	node.Next = node.Next.Next

	return head
}

func findNodeAt(head *ListNode, n int) *ListNode {
	node := head
	for i := 0; i < n; i++ {
		node = node.Next
	}

	return node
}
