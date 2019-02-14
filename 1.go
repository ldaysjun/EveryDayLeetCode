package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1},1))
}


func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i:=0;i<n-1;i++  {
		for j:=i+1;j<n;j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}

	return nil
}