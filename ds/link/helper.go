package link

func Slice2ListNodes(nums []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, num := range nums {
		node := &ListNode{
			Val: num,
		}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}
	return head
}
