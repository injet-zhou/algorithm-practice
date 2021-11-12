package link

type ListNode struct {
	Val int32
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	ptr := &node
	for (*ptr).Next != nil {
		(*ptr).Val = (*ptr).Next.Val
		if (*ptr).Next.Next == nil {
			(*ptr).Next = nil
			break
		} else {
			ptr = &(*ptr).Next
		}
	}
}

func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			dummy.Next = l2
			dummy = l2
			l2 = l2.Next
		} else {
			dummy.Next = l1
			dummy = l1
			l1 = l1.Next
		}
	}
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}
	return head.Next
}

func LastNNode(l *ListNode,n int) *ListNode {
	slow := &l
	fast := &l
	for i := 0; i < n; i++ {
		fast = &(*fast).Next
	}
	for (*fast).Next != nil {
		fast = &(*fast).Next
		slow = &(*slow).Next
	}
	return *slow
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	//h,t := &head,&(head.Next)
	return head
}