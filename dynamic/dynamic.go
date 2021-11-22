package dynamic

func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ { // 归纳法
			if nums[i]>nums[j] {
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
	}
	var res int
	for _, i := range dp {
		res = max(res,i)
	}
	return res
}

func max(a, b int) int {
	if a>= b {
		return a
	}
	return b
}
