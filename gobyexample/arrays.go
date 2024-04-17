package main

import "fmt"

func main() {

	var a [5]int // array creation
	fmt.Println("The created array is : ", a)

	a[4] = 5 //change value at 4th index
	fmt.Println("The created array is : ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("second array : ", b)
	fmt.Println("length of the second array is : ", len(b))

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	c := [...]int{1, 2, 3, 4, 5, 6} // compiler will count the number of elements for you.
	fmt.Println("third array is : ", c)

	// create a 2d array
	twoD := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoD)
}
