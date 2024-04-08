package main

import "fmt"

func main() {
	var size int
	fmt.Println("give the size of the array")
	fmt.Scanf("%d", &size)

	array := make([]int, size)

	fmt.Println("Enter ", size, "values separated by space: ")

	for i := 0; i < size; i++ {
		var num int
		fmt.Scan(&num)
		array[i] = num
	}

	fmt.Println("array created : ", array)

	var n_even int = 0
	for i := 0; i < size; i++ {
		if array[i]%2 == 0 {
			n_even++
		}
	}
	fmt.Println("The even numbers in the array is : ", n_even)
}
