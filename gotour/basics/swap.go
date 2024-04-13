// swap function example
package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x // returns the string values in a swaped order
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
