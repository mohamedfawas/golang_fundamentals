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

	// sorting
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if array[j] > array[i] {
				var temp int
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	fmt.Println("array after sorting in descending order is : ", array)
}
