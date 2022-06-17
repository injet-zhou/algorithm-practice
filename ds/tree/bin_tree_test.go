package tree

import (
	"testing"
)

func TestRevertTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	t.Log(isSameTree(root, q))
}

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	t.Log(preorderTraversal(root))
}

func TestGetPath(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	t.Log(pathSum(root, 22))
}

func TestIsUnivalTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		//Right: &TreeNode{
		//	Val: 1,
		//},
		//Left: &TreeNode{
		//	Val: 1,
		//	Left: &TreeNode{
		//		Val: 5,
		//	},
		//	Right: &TreeNode{
		//		Val: 1,
		//	},
		//},
	}
	t.Log(isUnivalTree(root))
}
