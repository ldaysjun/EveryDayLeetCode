package main

import "fmt"

func main() {
	testt := longestCommonPrefix([]string{"aa","a"})
	fmt.Println(testt)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	target := strs[0]
	for index,t := range target {
		flage := true
		for _,value := range strs {
			if len(value)<=index  {
				flage = false
				break
			}
			if string(value[index]) != string(t) {
				flage = false
				break
			}
		}
		if flage {
			res = res + string(t)
		}else {
			break
		}
	}

	return res
}
