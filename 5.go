package main

import (
	"fmt"
)

func main() {
	s := longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy")
	fmt.Println("result=",s)
}

func longestPalindrome(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	}
	var subStr = ""
	var result = string(s[0])
	for index,value := range s{
		subStr = string(value)
		subarray := s[index+1:]
		for i,_ := range subarray{
			subStr += string(subarray[:len(subarray)-i])
			if subStr[0:1] == subStr[len(subStr)-1:] {
				if judge(subStr) {
					if len(subStr)> len(result) {
						result = subStr
					}
					break
				}
			}
			subStr = string(value)
		}
	}

	return result
}

func judge(s string) bool{
	n := len(s)
	for i:=0;i<n;i++{
		m:=n-i-1
		if s[i] != s[m]  {
			return false
		}
		if i>n/2 {
			break
		}
	}
	return true
}


