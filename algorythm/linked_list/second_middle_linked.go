package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := getSize(head)
	middle := count / 2
	n := head
	for i := 0; i < middle; i++ {
		n = n.Next
	}

	return n
}

func getSize(head *ListNode) int {
	count := 0
	n := head

	for n != nil {
		count++
		n = n.Next
	}

	return count
}
