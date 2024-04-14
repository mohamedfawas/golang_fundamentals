package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println("read p : ", p)
	v.X = 5
	fmt.Println("read p : ", p)
}
