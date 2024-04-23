package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hai Fawas,Come to India please"
	split_str1 := strings.Split(str1, " ")
	fmt.Println(split_str1)
	fmt.Println(len(split_str1))
}
