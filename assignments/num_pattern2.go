package main

import "fmt"

var limit int = 4
var num int = 1

func main() {
	//fmt.Println()
	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println("")
	}
}
