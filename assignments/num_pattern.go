package main

import "fmt"

func main() {
	var num int = 5
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
