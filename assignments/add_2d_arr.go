package main

import "fmt"

func main() {
	var row, col int
	fmt.Println("Give the no: of rows in the input arrays: ")
	fmt.Scan(&row)
	fmt.Println("Give the no: of columns in the input arrays : ")
	fmt.Scan(&col)

	//initialize both arrays
	array1 := make([][]int, row)
	for i := range array1 {
		array1[i] = make([]int, col)
	}
	array2 := make([][]int, row)
	for i := range array2 {
		array2[i] = make([]int, col)
	}
	// add values to the arrays
	arr1, arr2 := getArray(array1, array2, row, col)
	sum_arr := addArray(arr1, arr2, row, col)
	displayArray(sum_arr)

}

func getArray(arr1 [][]int, arr2 [][]int, rows int, columns int) ([][]int, [][]int) {
	//Get values to the array
	fmt.Println("==========adding values in the first array============")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("Give the value to add at the position %d,%d :", i, j)
			fmt.Scan(&arr1[i][j])
		}
	}
	fmt.Println("==========adding values in the second array============")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("Give the value to add at the position %d,%d :", i, j)
			fmt.Scan(&arr2[i][j])
		}
	}
	return arr1, arr2
}

func addArray(arr1 [][]int, arr2 [][]int, rows int, columns int) [][]int {
	//Add array 1 and array 2

	//Initialize the array to add sum values.
	sum_array := make([][]int, rows)
	for i := range sum_array {
		sum_array[i] = make([]int, columns)
	}

	// sum
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			sum_array[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return sum_array
}

func displayArray(sum_arr [][]int) {
	//Display the array values
	fmt.Println("sum of both the arrays is : ", sum_arr)
}
