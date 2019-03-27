package main

import "fmt"

func main() {
	test := romanToInt("MMMCDXLII")
	fmt.Println(test)
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"M":1000,"CM":900,
		"D":500,"CD":400,
		"C":100,"XC":90,
		"L":50,"XL":40,
		"X":10,"IX":9,
		"V":5,"IV":4,"I":1,
	}
	total := 0
	index := 0
	for index < len(s)-1 {
		numD := romanMap[string(s[index]) + string(s[index+1])]
		if numD > 0 {
			total = numD + total
			index = index + 2
			continue
		}
		num := romanMap[string(s[index])]
		total = num + total
		index = index + 1

	}

	if index == len(s) -1 {
		total = total + romanMap[string(s[index])]
	}

	return total
}
