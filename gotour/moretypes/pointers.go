package main

import "fmt"

func main() {
	i := 42
	p := i
	fmt.Println("value of P : ", p)
	i = 55
	fmt.Println("value of P : ", p)
	fmt.Println("value of i : ", i)

	c := &i                                       // point to i
	fmt.Println("read i throught pointer : ", *c) // read i through the pointer

}
