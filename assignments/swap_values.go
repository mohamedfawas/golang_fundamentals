package main

import "fmt"

func swapValues(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	var x, y int
	fmt.Println("Give the first value :  ")
	fmt.Scan(&x)
	fmt.Println("Give the second value :  ")
	fmt.Scan(&y)

	x, y = swapValues(x, y)
	fmt.Printf("The values of the variables  x and y after swaping : %d, %d", x, y)
}
