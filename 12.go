package main

import "fmt"

func main() {
	x:=intToRoman(3442)
	fmt.Println(x)
}
func intToRoman(num int) string {
	numArray := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	romanArray := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	r := ""
	for index,v := range numArray{
		n := num / v
		for i:=0;i<n;i++{
			r = r + romanArray[index]
		}
		num = num - n * v
	}
	return r
}
