package day2

import (
	"bufio"
	"os"
	"strings"
)

func readReports(filepath string) (reports []Report) {
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
