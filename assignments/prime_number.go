package main

import "fmt"

func main() {
	fmt.Println("Give the input number : ")
	var num int
	fmt.Scan(&num)

	var count int
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			count++
			if (num / i) != i {
				count++
			}
		}
	}

	if count == 2 {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a prime number")
	}
}
