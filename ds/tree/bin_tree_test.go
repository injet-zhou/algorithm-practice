package tree

import "testing"

func TestRevertTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			//Left: &TreeNode{
			//	Val: 3,
			//},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	preorderTraverse(root)
}
