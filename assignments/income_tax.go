package main

import "fmt"

func incomeTaxCalculator(income float64) float64 {
	var tax_percentage float64
	if income <= 250000 {
		tax_percentage = 0
	} else if income > 250000 && income <= 500000 {
		tax_percentage = 0.05
	} else if income > 500000 && income <= 1000000 {
		tax_percentage = 0.2
	} else if income > 1000000 && income <= 5000000 {
		tax_percentage = 0.3
	}
	var income_tax float64 = income * tax_percentage
	return income_tax
}

func main() {

	fmt.Println("Please enter your income amount : ")
	var income float64
	fmt.Scan(&income)
	var income_tax float64 = incomeTaxCalculator(income)
	fmt.Printf("The income tax you have to pay is : %f ", income_tax)
}
