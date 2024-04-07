package main

import "fmt"

func main() {
	fmt.Println("Give an input number")
	var num int
	fmt.Scanf("%d", &num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*", i, " = ", num*i)
	}
}
