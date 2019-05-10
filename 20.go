package main

import "fmt"

func main() {
	test := "()[]{}"
	for _,v := range test{
		fmt.Println(v)
	}


	fmt.Println(string(124))
}
