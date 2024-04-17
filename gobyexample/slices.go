package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make(type, 0)

	fmt.Println("test")
}
