package main

import "fmt"

func main() {
	result := convert("PAYPALISHIRING",3)
	fmt.Println(result)

}

func convert(s string, numRows int) string {

	if len(s) == 0 || numRows == 0{
		return ""
	}
	if numRows == 1 {
		return s
	}
	//字符串数
	strNum := len(s)
	//分组书
	groupNum := numRows + numRows - 2
	//分组列数
	groupCol := groupNum - numRows + 1
	//列数
	col := (strNum / groupNum) * groupCol

	//求剩余需要几列，只可能是1列或2列
	last := strNum % groupNum
	if last > 0 {
		col += last % numRows + 1
	}

	//遍历二维数组，映射到s的位置公司 s[n] = s[j*2+i]
	row := numRows
	result := ""
	for i:=0;i<row ;i++  {
		for j:=0;j<col;j++  {
			if j % groupCol == 0 {
				if j*2+i >= len(s) {
					continue
				}
				result += string(s[j*2+i])
			}else {
				//计算该位置是否应该有值
				if i == row-j%groupCol-1{
					if j*2+i >= len(s) {
						continue
					}
					result += string(s[j*2+i])
				}
			}
		}
	}
	return result
}