package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	count := 1
	for i:=0;i<len(nums)-1;i++ {
		if nums[i]<nums[i+1]{
			count ++
		}else {
			count = 1
		}

		if count>result {
			result = count
		}
	}
	return result
}


//func findLengthOfLCIS(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	res := 0
//	dp := make([]int,n)
//	for i:=0;i<n;i++ {
//		dp[i]=1
//		if i>0&&nums[i]>nums[i-1] {
//			dp[i]=dp[i-1]+1
//		}
//		if res<dp[i] {
//			res=dp[i]
//		}
//	}
//	return res
//}
