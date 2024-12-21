package day2

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"math"
)

type Report struct {
	Levels []int
}

func createReport(s []string) Report {
	var levels []int

	for _, item := range s {
		levels = append(levels, helper.ExtractInt(item))
	}

	return Report{Levels: levels}
}

/*
A report is only considered safe if both of the following are true:
1) The levels are either all increasing or all descreasing
2) Any two adjacent levels differ by at least one and at most three.
*/
func (r *Report) isReportSafe() bool {
	safe := true

	if len(r.Levels) == 1 {
		return safe
	}

	for i := 1; i < len(r.Levels)-1; i++ {
		left_comparison := math.Abs(float64(r.Levels[i] - r.Levels[i-1]))
		if left_comparison < 1.0 || left_comparison > 3.0 {
			safe = false
			break
		}

		right_comparison := math.Abs(float64(r.Levels[i] - r.Levels[i+1]))
		if right_comparison < 1.0 || right_comparison > 3.0 {
			safe = false
			break
		}

		// Catch Peak
		if r.Levels[i-1] <= r.Levels[i] && r.Levels[i] >= r.Levels[i+1] {
			safe = false
			break
		}

		// Catch Trough
		if r.Levels[i-1] >= r.Levels[i] && r.Levels[i] <= r.Levels[i+1] {
			safe = false
			break
		}
	}

	return safe
}

func (r *Report) canDampenerMakeSafe() bool {
	dampener_made_safe := false

	for i := 0; i < len(r.Levels); i++ {
		fresh_slice := make([]int, len(r.Levels))
		copy(fresh_slice, r.Levels)
		all_bar_one := Report{Levels: deleteElement(fresh_slice, i)}
		if all_bar_one.isReportSafe() {
			dampener_made_safe = true
			break
		}
	}

	return dampener_made_safe
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
