package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Please input your total marks")
	fmt.Scanf("%d", &marks)

	if marks >= 0 && marks < 50 {
		fmt.Println("Failed")
	} else if marks >= 50 && marks < 60 {
		fmt.Println("E")
	} else if marks >= 60 && marks < 70 {
		fmt.Println("D")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("C")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("B")
	} else if marks >= 90 && marks <= 100 {
		fmt.Println("A")
	} else {
		fmt.Println("Invalid input")
	}
}
