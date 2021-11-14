package link

import (
	"fmt"
	"math/rand"
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
		Val: 4,
	}
	nn := &ListNode{
		Val: 5,
	}
	first.Next = n5
	n5.Next = n1
	n1.Next = n9
	n9.Next = nn
	n := middleNode(first)
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

func TestBitMove(t *testing.T) {
	first := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 0,
	}
	n3 := &ListNode{
		Val: 1,
	}
	n4 := &ListNode{
		Val: 1,
	}
	n5 := &ListNode{
		Val: 0,
	}
	first.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	var a = 0
	cur := first
	for cur != nil {
		a = a << 1
		a = a | cur.Val
		t.Log(a)
		cur = cur.Next

	}
}

func TestIsEvenOrOdd(t *testing.T) {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}

}

func TestGetDecimalValue(t *testing.T) {
	first := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 0,
	}
	n3 := &ListNode{
		Val: 1,
	}
	n4 := &ListNode{
		Val: 1,
	}
	n5 := &ListNode{
		Val: 0,
	}
	first.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	t.Log(getDecimalValue(&ListNode{Val: 0}))
}

func TestRemoveElements(t *testing.T) {
	first := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 3,
	}
	n3 := &ListNode{
		Val: 2,
	}
	n4 := &ListNode{
		Val: 4,
	}
	n5 := &ListNode{
		Val: 2,
	}
	first.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n := removeElements(first,2)
	for n != nil {
		t.Log(n.Val)
		n = n.Next
	}
}