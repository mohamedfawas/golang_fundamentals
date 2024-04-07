package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Please input your total marks")
	fmt.Scanf("%d", &marks)

	if marks >= 0 && marks < 50 {
		fmt.Println("failed")
	} else if marks >= 50 && marks <= 100 {
		fmt.Println("passed")
	} else {
		fmt.Println("invalid input is given")
	}
}
