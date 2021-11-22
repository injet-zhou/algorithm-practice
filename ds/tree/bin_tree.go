package tree

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTreeNode(root)
	return root
}
func invertTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := &TreeNode{}
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTreeNode(root.Left)
	invertTreeNode(root.Right)
}

func TraverseTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	TraverseTree(root.Left)
	TraverseTree(root.Right)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left,root.Right)
	return root
}

func connectTwoNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNode(node1.Left,node1.Right)
	connectTwoNode(node2.Left,node2.Right)
	connectTwoNode(node1.Right,node2.Left)
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	inorderTraversalNode(root,&nums)
	return nums
}
func inorderTraversalNode(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorderTraversalNode(root.Left,nums)
	*nums = append(*nums,root.Val)
	inorderTraversalNode(root.Right,nums)
}

func preorderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val," ")
	preorderTraverse(root.Left)
	preorderTraverse(root.Right)
}
// 二叉树展开为链表
func flatten(root *TreeNode)  {
	flattenTree(root)
}

func flattenTree(root *TreeNode) {
	if root == nil {
		return
	}
	flattenTree(root.Left)
	flattenTree(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}