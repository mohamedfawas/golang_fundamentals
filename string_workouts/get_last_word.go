package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hai I am fawas"
	split_str1 := strings.Split(str1, " ")
	fmt.Println(split_str1[len(split_str1)-1])
}
