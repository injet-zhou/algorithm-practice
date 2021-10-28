package huisu

func Combine(n int, k int) [][]int {
	var (
		route [][]int
		track []int
		count int
	)
	route = make([][]int,GetCombineCount(n,k))
	BackTrack2(k,n,track,route,&count)
	return route
}

func BackTrack2(k int,n int,track []int,route [][]int,count *int) {
	if len(track) == k {
		route[*count] = make([]int,len(track))
		copy(route[*count],track)
		*count++
		return
	}
	for i := 1; i < n+1; i++ {
		if Contain(track,i) {
			continue
		}
		if len(track) != 0 && track[len(track)-1] > i {
			continue
		}
		track = append(track,i)
		BackTrack2(k,n,track,route,count)
		track = track[:len(track)-1]
	}
}

func GetCombineCount(n, k int) int {
	res := Fact(n) / (Fact(k)*Fact(n-k))
	return res
}
