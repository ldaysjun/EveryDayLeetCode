package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	result := 0
	y := 1

	//负数转整数计算
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}
	for true {
		//分解数据
		r := x / y % 10
		//结果为0 可能是超出大小，结束循环，判断是否大于y
		if r == 0 && x < y{
			break
		}
		//累加
		result = result * 10 + r
		y = y * 10
	}
	if isNeg {
		result = -result
	}
	//计算是否超出范围
	if result>2147483647 || result< -2147483648{
		result = 0
	}
	return result
}