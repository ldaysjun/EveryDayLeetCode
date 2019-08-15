package main

func main() {

}


func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	max := 0
	min := prices[0]

	for i:=0;i<n;i++ {
		tempMax := prices[i] - min
		if max < tempMax{
			max = tempMax
		}

		if min > prices[i] {
			min = prices[i]
		}
	}

	return max
}