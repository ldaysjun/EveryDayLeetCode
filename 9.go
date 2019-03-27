package main

func main() {
	isPalindrome(121)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y:=1
	result := 0
	for true{
		r := x/y%10
		if r == 0 && x<y {
			break
		}
		result = result*10+r
		y = y * 10
	}
	if result == x{
		return true
	}

	return false
}
