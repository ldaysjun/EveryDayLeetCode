package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	test := threeSum([]int{-1,0,1,2,-1,-4})
	fmt.Println(test)
}


func threeSum(nums []int) [][]int {
	result := make([][]int,0)
	for i:=0;i<len(nums) ;i++  {
		target := -nums[i]
		for j:=i+1; j<len(nums)-1; j++ {
			for k:=j+1;k<len(nums);k++  {
				a := nums[i]
				b := nums[j]
				c := nums[k]
				if nums[k] + nums[j]  == target {
					isex := false
					temp := []int{a,b,c}
					sort.Ints(temp)
					for	_,value := range result{
						if reflect.DeepEqual(value,temp) {
							isex = true
						}
					}
					if !isex {
						result = append(result, temp)
					}
				}
			}
		}
	}
	return result
}
