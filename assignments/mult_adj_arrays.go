package main

import "fmt"

func main() {
	//array input
	var size int
	fmt.Println("Give the size of the array: ")
	fmt.Scan(&size)
	array := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Give the number to add at the position %d : ", i)
		fmt.Scan(&array[i])
	}
	fmt.Println("The created array is :", array)

	//array to store multiplied values.
	mult_array := make([]int, size-1)
	for i := 0; i < size; i++ {
		if i == size-1 {
			continue
		} else {
			mult_array[i] = array[i] * array[i+1]
		}
	}
	fmt.Println("The array after multiplication is :", mult_array)
}
