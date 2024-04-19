package main

import "fmt"

type Calculator struct{} // calculator struct to keep the methods

func (c Calculator) addition(num1, num2 float64) float64 {
	return num1 + num2
}
func (c Calculator) subtraction(num1, num2 float64) float64 {
	return num1 - num2
}
func (c Calculator) multiplication(num1, num2 float64) float64 {
	return num1 * num2
}
func (c Calculator) division(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

func main() {
	var operator int
	var num1, num2 float64
	fmt.Println("Give the first input number num1 : ")
	fmt.Scan(&num1)
	fmt.Println("Give the first input number num2 : ")
	fmt.Scan(&num2)
	calc := Calculator{}

	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&operator)

	switch operator {
	case 1:
		fmt.Println(calc.addition(num1, num2))
	case 2:
		fmt.Println(calc.subtraction(num1, num2))
	case 3:
		fmt.Println(calc.multiplication(num1, num2))
	case 4:
		fmt.Println(calc.division(num1, num2))
	default:
		fmt.Println("Invalid input")

	}
}
