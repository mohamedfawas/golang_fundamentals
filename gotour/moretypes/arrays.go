package main

import "fmt"

func main() {
	var a [2]string
	fmt.Println("The created array is : ", a)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("The created array is : ", a)
}
