package huisu



type TreeNode struct {
   Val int
    Left *TreeNode
    Right *TreeNode
}


func CombineSum(root *TreeNode, target int) [][]int {
	var (
		route [][]int
		track []int
		count int
	)
	route = make([][]int,50)
	BackTrackSum(root,track,route ,&count,target)
	return route
}

func BackTrackSum(root *TreeNode,track []int,route [][]int,count *int,target int) {
	if root != nil {
		track = append(track,root.Val)
		if SumIntSlice(track) == target {
			route[*count] = make([]int,len(track))
			copy(route[*count],track)
			*count++
			return
		}
		BackTrackSum(root.Left,track,route ,count,target)
		BackTrackSum(root.Right,track,route ,count,target)
		track = track[:len(track)-1]
	}
}

func SumIntSlice(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
