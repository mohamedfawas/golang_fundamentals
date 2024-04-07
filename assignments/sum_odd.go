package main

import "fmt"

func main() {
	fmt.Print("Give the input number: ")
	var num int
	fmt.Scanf("%d", &num)

	var sum int = 0
	for i := 0; i <= num; i++ {
		if i%2 == 1 {
			sum += i
		}
	}
	fmt.Println("The sum of odd numbers is : ", sum)
}
