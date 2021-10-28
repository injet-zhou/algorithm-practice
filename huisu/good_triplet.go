package huisu

//func countGoodTriplets(arr []int, a int, b int, c int) int {
//
//}

func BackTrack3(num []int,track []int,count *int) {
	if len(num) == len(track) {
		//route[*count] = make([]int,len(track))
		//copy(route[*count],track)
		*count++
		return
	}
	for i := 0; i < len(num); i++ {
		if Contain(track,num[i]) {
			continue
		}
		track = append(track,num[i])
		BackTrack3(num,track,count)
		track = track[:len(track)-1]
	}
}

