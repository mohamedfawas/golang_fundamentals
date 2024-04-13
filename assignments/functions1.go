package main

import "fmt"

func main() {
	//take array input
	var size int
	fmt.Println("Give the size of the input array :")
	fmt.Scan(&size)

	//declare an array
	array := make([]int, size)

	//Input values in the array
	getArray(array, size)

	// Display the created array
	displayArray(array)
}

func getArray(array []int, size int) []int {
	for i := 0; i < size; i++ {
		fmt.Printf("Give the value to add at the position %d : ", i)
		fmt.Scan(&array[i])
	}
	return array
}

func displayArray(array []int) []int {
	fmt.Println("The created array is :", array)
	return array
}
