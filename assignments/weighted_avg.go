package main

import "fmt"

func wtAverageCalculator(written_test int, lab_exam int, assignments int) float64 {
	var written_wt float64 = 0.7
	var lab_wt float64 = 0.2
	var assign_wt float64 = 0.1

	var grade float64 = float64(written_test)*written_wt + float64(lab_exam)*lab_wt + float64(assignments)*assign_wt
	return grade
}

func main() {
	var wrt_test, lab_ex, assg int
	fmt.Println("Give the marks obtained in written test: ")
	fmt.Scan(&wrt_test)
	fmt.Println("Give the marks obtained in Lab exam: ")
	fmt.Scan(&lab_ex)
	fmt.Println("Give the marks obtained in assignments : ")
	fmt.Scan(&assg)

	var final_grade float64
	final_grade = wtAverageCalculator(wrt_test, lab_ex, assg)
	fmt.Println("The final grade obtained by the student is : ", final_grade)
}
