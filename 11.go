package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
//
//func maxArea(height []int) int {
//	result := 0
//	for index,value := range height{
//		for n,val := range height[index+1:]{
//			w := n + 1
//			h := 0
//			if val>=value{
//				h = value
//			}else {
//				h = val
//			}
//			r := w * h
//			if result < r{
//				result = r
//			}
//		}
//	}
//
//	return result
//}


func maxArea(height []int) int {
	result := 0
	l := 0
	r := len(height) - 1
	w := len(height) - 1
	for l<r {
		h := 0
		if  height[l]>height[r]{
			h = height[r]
			r = r - 1
		}else {
			h = height[l]
			l = l + 1
		}

		are := h * w
		if result < are{
			result = are
		}
		w = w - 1
	}
	return result
}
