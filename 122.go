package main

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	
	total := 0
	for i:=0;i< len(prices) - 1;i++  {
		temp := prices[i+1] - prices[i]
		if temp > 0 {
			total += temp
		}
	}

	return total

}
