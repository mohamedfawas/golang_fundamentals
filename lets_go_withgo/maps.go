package main

import "fmt"

func main() {
	fmt.Println("Learn maps")
	languages := make(map[string]string) //[key]value
	languages["JS"] = "javascript"
	languages["CP"] = "C++"
	languages["PY"] = "Python"
	fmt.Println("List of all languages : ", languages) //They're not comma separated in golang
	fmt.Println("JS is the short for :", languages["JS"])

	delete(languages, "PY")
	fmt.Println("======after removing RUBY======================")
	fmt.Println("List of all languages : ", languages)

	//loops
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}
}
