package bfs

import (
	"sinkcode.cn/alg/ds/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	t.Log(BFS(root))

}

func BFS(root *TreeNode) (step int32) {
	if root == nil {
		return step
	}
	q := queue.New()
	q.Add(root)
	// root也算深度
	step++
	for q.Length() != 0 {
		size := q.Length()
		for i := 0; i < size; i++ {
			// 获取队首，但是没有删除
			cur := q.Peek().(*TreeNode)
			// 记得删除
			q.Remove()
			if cur.Left == nil && cur.Right == nil {
				return step
			}
			// 如果右子树不为空则加入队列
			if cur.Right != nil {
				q.Add(cur.Right)
			}
			// 如果左子树不为空则加入队列
			if cur.Left != nil {
				q.Add(cur.Left)
			}
		}
		step++
	}
	return
}
