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
	v3 := m["k3"]
	fmt.Println("value v2 : ", v3)

	fmt.Println("length of m is ", len(m))

	delete(m, "k2")
	fmt.Println("the map after deleting the value k2 : ", m)

	_, prs := m["k1"] // returns true or false, which means value is present or not
	fmt.Println("prs:", prs)
	//fmt.Println(m["k1"])
}
