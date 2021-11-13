package link

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	first := &ListNode{
		Val: 1,
	}
	n5 := &ListNode{
		Val: 2,
	}
	n1 := &ListNode{
		Val: 3,
	}
	n9 := &ListNode{
		Val: 3,
	}
	first.Next = n5
	n5.Next = n1
	n1.Next = n9
	n := deleteDuplicatesV2(first)
	for n != nil {
		t.Log(n.Val)
		n = n.Next
	}
}

func TestMergeTwoList(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2:= &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := MergeTwoList(l1,l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}