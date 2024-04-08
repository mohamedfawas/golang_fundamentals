package main

import "fmt"

func main() {
	var size int
	fmt.Println("give the size of the input array")
	fmt.Scan(&size)5
	

	array := make([]int, size)
	fmt.Println("enter ", size, "values separated by spaces")

	for i := 0; i < size; i++ {
		var num int
		fmt.Scan(&num)
		array[i] = num
	}

	fmt.Println("array created : ", array)
}
