package main

import "sort"

func findKthLargest(nums []int, k int) int {
	for i:=0;i< len(nums)-1;i++  {
		for j:=0;j<len(nums)-i-1 ;j++  {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	result := nums[len(nums)-k]
	return result
}



func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	result := nums[len(nums)-k]
	return result
}