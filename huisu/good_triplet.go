package huisu

import "sinkcode.cn/alg/util"

func CountGoodTriplets(arr []int, a int, b int, c int) int {
	s := &GoodService{
		A: a,
		B: b,
		C: c,
	}
	var count int = 0
	var track []int
	s.BackTrack3(arr,track,&count)
	return count
}

type GoodService struct {
	A int
	B int
	C int
}

func (s *GoodService)BackTrack3(num []int,track []int,count *int) {
	if len(track) == 3 && s.IsEnd(track) {
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
		s.BackTrack3(num,track,count)
		track = track[:len(track)-1]
	}
}

func (s *GoodService) IsEnd(track []int) bool {
	a := util.Abs(track[0] - track[1]) <= s.A
	b := util.Abs(track[1] - track[2]) <= s.B
	c := util.Abs(track[0] - track[2]) <= s.C
	return a && b && c
}

