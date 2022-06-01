package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecondMiddle(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	n := middleNode(h)
	assert.Equal(t, 3, n.Val)

	h = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{Val: 6},
					},
				},
			},
		},
	}
	n = middleNode(h)
	assert.Equal(t, 4, n.Val)
}

func TestFindAt(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{Val: 6},
					},
				},
			},
		},
	}

	n := findNodeAt(h, 2)
	assert.Equal(t, 3, n.Val)
	m := findNodeAt(h, 0)
	assert.Equal(t, 1, m.Val)
}

func TestRemoveNFromEnd(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	n := removeNthFromEnd(h, 1)
	assert.Equal(t, 1, n.Val)

	h = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	m := removeNthFromEnd(h, 2)
	assert.Equal(t, 2, m.Val)
}
