package main

import "fmt"

func main() {
	primes := [7]int{2, 3, 5, 7, 11, 13, 17}
	var s []int = primes[1:4]
	fmt.Println(s)
}
