package huisu

import "testing"

func TestFact(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 2,
					},
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
	t.Log(CombineSum(tree,22))
}

func TestCountGoodTriplets(t *testing.T) {
	arr := []int{3,0,1,1,9,7}
	count := CountGoodTriplets(arr,7,2,3)
	t.Log(count)
}

