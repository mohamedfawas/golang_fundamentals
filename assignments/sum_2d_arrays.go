package main

import "fmt"

func main() {
	//fmt.Println("creation of first array")
	var row, col int
	fmt.Println("Give the number of rows in the input array: ")
	fmt.Scan(&row)
	println("Give the number of columns in the input array: ")
	fmt.Scan(&col)
	//fmt.Println("test print: ", row, " ", col)
	// Intialise arrays with zeros
	array1 := make([][]int, row)
	for i := range array1 {
		array1[i] = make([]int, col)
	}
	//fmt.Println(" the created array is : ", array1)
	fmt.Println("Give values to add in the first array: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Give the value to add at the position %d,%d :", i, j)
			fmt.Scan(&array1[i][j])
		}
	}
	//fmt.Println(" the created array is : ", array1)

	// create second array
	array2 := make([][]int, row)
	for i := range array2 {
		array2[i] = make([]int, col)
	}
	//fmt.Println(" the created array is : ", array2)
	fmt.Println("Give values to add in the second array: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Give the value to add at the position %d,%d :", i, j)
			fmt.Scan(&array2[i][j])
		}
	}
	fmt.Println(" the created first array is : ", array1)
	fmt.Println(" the created second array is : ", array2)

	// create second array
	sum_array := make([][]int, row)
	for i := range sum_array {
		sum_array[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum_array[i][j] = array1[i][j] + array2[i][j]
		}
	}
	fmt.Println(" the array created after adding values in given two arrays : ", sum_array)

}
