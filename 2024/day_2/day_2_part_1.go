package day2

import (
	"fmt"
)

func SolveDay2Part1() {
	reports := readReports("day_2_input.txt")

	count := 0

	for _, report := range reports {
		if report.isReportSafe() {
			count += 1
		}
	}

	fmt.Println(count)
}
