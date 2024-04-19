package main

import "fmt"

type twoDArray struct {
	rows  int
	cols  int
	array [][]int
}

func createTwoDArray(rows, cols int) *twoDArray {
	return &twoDArray{
		rows:  rows,
		cols:  cols,
		array: make([][]int, rows),
	}
}

func (t *twoDArray) getArray() {
	fmt.Println("enter  the values of the array : ")
	for i := 0; i < t.rows; i++ {
		t.array[i] = make([]int, t.cols)
		for j := 0; j < t.cols; j++ {
			fmt.Printf("Enter the value to add at the postion %d,%d  : ", i, j)
			fmt.Scan(&t.array[i][j])

		}
	}
}

func (t *twoDArray) displayArray() {
	fmt.Println("Array values are : ")
	for _, row := range t.array {
		fmt.Println(row)
	}
}

func main() {
	//fmt.Println("test")
	var rows, cols int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&rows)
	fmt.Print("Enter number of columns: ")
	fmt.Scanln(&cols)

	array := createTwoDArray(rows, cols)
	array.getArray()
	array.displayArray()
}
