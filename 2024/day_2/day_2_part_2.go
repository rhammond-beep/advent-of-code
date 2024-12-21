package day2

import (
	"fmt"
)

func SolveDay2Part2() {
	reports := readReports("day_2_input.txt")
	count := 0

	for _, report := range reports {
		if report.isReportSafe() || report.canDampenerMakeSafe() {
			count++
		}
	}
	fmt.Println(count)
}
