package main

import (
	"fmt"
	//"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println("the created map is : ", m)

	v1 := m["k1"]
	fmt.Println("value v1 : ", v1)
	v2 := m["k1"]
	fmt.Println("value v2 : ", v2)

	fmt.Println("length of m is ", len(m))
}
