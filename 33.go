package main

import "fmt"

func main() {
	//test := search([]int{0,1,2,3,4,5},3)
	var te string

	test2 := "你还"
	for _,v := range test2{
		fmt.Println(string(v))
		fmt.Println(v)
	}


	//fmt.Println(test)
}

func search(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	for start<=end {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}
		mid := (start + end)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[start]{
			if nums[mid]>target && nums[start]<target{
				end = mid - 1
			}else {
				start = mid + 1
			}
		}else {
			if nums[mid]<target && nums[end]>target {
				start = mid + 1
			}else {
				end = mid - 1
			}
		}



	}
	return -1
}