package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Println("Give the first string : ")
	fmt.Scan(&str1)
	fmt.Println("Give the second string : ")
	fmt.Scan(&str2)

	fmt.Println("result of comparing the given strings : ", strings.Compare(str1, str2))
}

/*
Return 0, if str1 == str2.
Return 1, if str1 > str2.
Return -1, if str1 < str2.
*/
