package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func (r *Report) IsReportSafe() bool {
	safe := true

	for i := 1; i < len(r.levels)-1; i++ {
		left_comparison := math.Abs(float64(r.levels[i] - r.levels[i-1]))
		if left_comparison < 1 || left_comparison > 3 {
			safe = false
			break
		}

		right_comparison := math.Abs(float64(r.levels[i] - r.levels[i+1]))
		if right_comparison < 1 || right_comparison > 3 {
			safe = false
			break
		}
	}

	return safe
}

/*
Checking the safety for each given report
*/
func main() {
	reports := ReadChallengeInput("../day_2_input.txt")

	count := 0

	for _, report := range reports {
		if report.IsReportSafe() {
			count += 1
		}
	}

	fmt.Println(count)
}

func ReadChallengeInput(filepath string) (reports []Report) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), " ")
		reports = append(reports, createReport(pair))
	}
	return
}

func createReport(s []string) Report {
	var levels []int

	for _, item := range s {
		levels = append(levels, extractInt(item))
	}

	return Report{levels: levels}
}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}