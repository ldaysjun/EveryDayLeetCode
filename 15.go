package main

import (
	"fmt"
	"sort"
)

//func main() {
//	test := threeSum([]int{-1,0,1,2,-1,-4})
//	fmt.Println(test)
//
//
//
//
//}


func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result :=  make([][]int,0)
	sort.Ints(nums)
	for i:=0;i<len(nums)-2;i++ {
		l := i + 1
		r := len(nums)-1
		target := -nums[i]
		for l<r{
			a := nums[l]
			b := nums[r]
			if a+b == target {
				temp := []int{a,b,-target}
				isexit := false
				for _,value := range result{
					if value[0] == temp[0] && value[1] == temp[1] {
						isexit = true
						break
					}
				}
				if !isexit {
					result = append(result, temp)
				}
				r --
				l ++
			}else if a + b < target {
				l ++ 
			}else {
				r --
			}
		}
	}
	return result
}
