package main

import "fmt"

func main() {
	var row int
	var col int
	fmt.Println("Give the no: of rows in the input array")
	fmt.Scan(&row)
	fmt.Println("Give the no: of columns in the input array")
	fmt.Scan(&col)

	array1 := make([][]int, row)
	for i := range array1 {
		array1[i] = make([]int, col)
	}

	//fmt.Println(array1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Give the number to add at the position [%d],[%d] : ", i, j)
			fmt.Scan(&array1[i][j])
		}
	}

	fmt.Println(array1)
}
