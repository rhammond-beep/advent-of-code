package day2

import (
	"fmt"
)

/*
Checking the safety for each given report
*/
func SolveDay2Part1() {
	reports := readReports("../day_2_input.txt")

	count := 0

	for _, report := range reports {
		if report.isReportSafe() {
			count += 1
		}
	}

	fmt.Println(count)
}
