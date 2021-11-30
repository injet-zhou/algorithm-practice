package tree

import (
	"fmt"
)

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	return traverseTwoTree(p,q) // 同步递归遍历
}

func traverseTwoTree(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if (q == nil && p != nil ) || (q != nil && p == nil) {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return traverseTwoTree(p.Left,q.Left) && traverseTwoTree(p.Right,q.Right)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traverseBinTree(root,0,targetSum)
}

func traverseBinTree(root *TreeNode, sum int, target int) bool {
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		if target == sum {
			return true
		}
	}
	left := false
	right := false
	if root.Left != nil {
		left =  traverseBinTree(root.Left,sum,target)
	}
	if root.Right != nil {
		right = traverseBinTree(root.Right,sum,target)
	}
	return left || right
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)+1
	right := maxDepth(root.Right)+1
	if left >= right {
		return left
	} else {
		return right
	}
}

func preorderTraversal(root *TreeNode) []int {
	nums := &[]int{}
	preTraverse(root,nums)
	return *nums
}
func preTraverse(root *TreeNode,nodes *[]int) {
	if root == nil {
		return
	}
	*nodes = append(*nodes,root.Val)
	preTraverse(root.Left,nodes)
	preTraverse(root.Right,nodes)
}
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return sumLeft(root)
}

func sumLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			return root.Left.Val + sumLeft(root.Right)
		}
	}
	left := sumLeft(root.Left)
	right := sumLeft(root.Right)
	return left + right
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	route := [][]int{}
	path := []int{}
	sums := 0
	getPathSum(root,path,targetSum,&route,sums)
	return route
}

func isLeaveNode(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

func getPathSum(root *TreeNode, path []int, target int,route *[][]int,sums int) {
	if root == nil {
		return
	}
	path = append(path,root.Val)
	sums += root.Val
	if isLeaveNode(root) {
		fmt.Println("path: ",path)
		fmt.Println("sum: ",sums," target: ",target)
		if sums == target {
			*route = append(*route,copySlice(path))
		}
	}
	getPathSum(root.Left,path,target,route,sums)
	getPathSum(root.Right,path,target,route,sums)
	path = path[0:len(path)-1]
	sums -= root.Val
}

func copySlice(path []int) (sum []int) {
	sum = make([]int,0,len(path))
	sum = append(sum,path...)
	return sum
}