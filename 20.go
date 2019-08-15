package main

import (
	"fmt"
)


//俄罗斯方块

func main() {
	//test := "()[]{}"
	//for i:=0;i< len(test);  i++{
	//	fmt.Println(test[i])
	//}

	//test := "(([]){})"
	//total := test[:2] + test[4:]
	//fmt.Println(total)

	if isValid("((") {
		fmt.Println("yes")
	}else {
		fmt.Println("NO")
	}
}

func isValid(s string) bool {
	if len(s) % 2 == 1 {
		return false
	}

	l := 0

	for len(s)>0 {
		first := s[l]
		if first == 40 && l+1<len(s){
			if s[l+1] == 41 {
				s = s[:l] + s[l+2:]
				l = 0
			}else {
				l++
			}
		}else if first == 91 && l+1<len(s){
			if s[l+1] == 93 {
				s = s[:l] + s[l+2:]
				l = 0
			}else {
				l++
			}
		}else if first == 123 && l+1<len(s){
			if s[l+1] == 125 {
				s = s[:l] + s[l+2:]
				l = 0
			}else {
				l++
			}
		}else if l == len(s){
			if l>0 {
				return false
			}else {
				return true
			}
		}else {
			return false
		}
	}
	return true
}