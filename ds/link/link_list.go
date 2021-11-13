package link

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	ptr := &node
	for (*ptr) != nil {
		if (*ptr).Next != nil {
			(*ptr).Val = (*ptr).Next.Val
		}
		if (*ptr).Next == nil {
			*ptr = nil
			break
		} else {
			ptr = &(*ptr).Next
		}
	}
	fmt.Printf("pointer: %p",*ptr)
}

// MergeTwoList 拉链法合并两个有序链表
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
	for (*fast) != nil {
		fast = &(*fast).Next
		slow = &(*slow).Next
	}
	return *slow
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &head
	t := &(head.Next)
	loop:
	for t != nil {
		if (*h).Val == (*t).Val {
			(*h).Next = (*t).Next
			if *t == nil {
				break loop
			}
		} else {
			if (*t).Next != nil {
				h = t
				t = &(*t).Next
			} else {
				t = nil
			}
		}
	}
	return head
	// 更优版本
	//if head == nil {
	//	return nil
	//}
	//
	//cur := head
	//for cur.Next != nil {
	//	if cur.Val == cur.Next.Val {
	//		cur.Next = cur.Next.Next
	//	} else {
	//		cur = cur.Next
	//	}
	//}
	//
	//return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := &head
	fast := &head
	// fast先移动 n步
	for i := 0; i < n; i++ {
		fast = &(*fast).Next
	}
	// 两个指针同步移动
	for (*fast) != nil {
		fast = &(*fast).Next
		slow = &(*slow).Next
	}
	// 更好的方法：先找到倒数n+1个，然后进行删除：(*slow).Next = (*slow).Next.Next
	for (*slow) != nil {
		if (*slow).Next != nil {
			(*slow).Val = (*slow).Next.Val
		}
		if (*slow).Next == nil {
			*slow = nil
			break
		} else {
			slow = &(*slow).Next
		}
	}
	return head
}

func deleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	h := dummy // 虚拟头节点
	dummy.Next = head
	cur := head
	val := -200
	for cur != nil {
		// 加一个判断防止cur.next为空时调用cur.next.val panic
		if cur.Next != nil {
			if cur.Val == cur.Next.Val {
				val = cur.Val // 记录删除值
				cur.Next = cur.Next.Next // 删除重复节点
			}
		}
		// 如果上面进行了删除同时当前的节点与上面删除的节点相同，则删除当前节点
		if val == cur.Val {
			// 删除cur
			dummy.Next = cur.Next
			cur = cur.Next
		} else {
			// 否则往后移动
			dummy = cur
			cur = cur.Next
		}
	}
	return h.Next
}