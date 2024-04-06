package main

import "fmt"

func main() {
	var num1 int
	var num2 float64
	var sum float64

	fmt.Println("Give the first and second number inputs : ")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%f", &num2)

	sum = float64(num1) + num2
	fmt.Println("Sum = ", sum)
}
