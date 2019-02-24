package main

import "fmt"

func main() {
	r := lengthOfLongestSubstring("dvdf")
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	array:=make([]int32,0)
	var result = 0
	for _,value := range s{
		array = append(array, value)
		for i,v := range array[:len(array)-1]{
			if value == v{
				array = array[i+1:]
			}
		}
		if result < len(array){
			result = len(array)
		}
	}

	return result
}
