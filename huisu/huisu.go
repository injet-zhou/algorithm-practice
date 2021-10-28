package huisu

func FullSort(num []int) [][]int {
	var (
		route [][]int
		track []int
		count int
	)
	route = make([][]int,Fact(len(num)))
	BackTrack(num,track,route,&count)
	return route
}

func BackTrack(num []int,track []int,route [][]int,count *int) {
	if len(num) == len(track) {
		route[*count] = make([]int,len(track))
		copy(route[*count],track)
		*count++
		return
	}
	for i := 0; i < len(num); i++ {
		if Contain(track,num[i]) {
			continue
		}
		track = append(track,num[i])
		BackTrack(num,track,route,count)
		track = track[:len(track)-1]
	}
}



func Contain(track []int,target int) bool {
	if len(track) == 0 {
		return false
	}
	for _, i := range track {
		if i == target {
			return true
		}
	}
	return false
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n*Fact(n-1)
}
