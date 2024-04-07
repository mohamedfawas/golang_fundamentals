package main

import "fmt"

func main() {
	var P int
	var R float64
	var n float64
	var SI float64

	fmt.Println("Give the Principal amount, Interest rate and no: of years :")
	fmt.Scanf("%d", &P)
	fmt.Scanf("%f", &R)
	fmt.Scanf("%f", &n)

	SI = (float64(P) * R * n) / 100
	fmt.Println("The calculated simple interest is :", SI)

}
