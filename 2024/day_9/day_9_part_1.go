package day9

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

var (
	file_num, free_space int
)

func SolveDay9Part1() {

	puzzleInput := "2333133121414131402"

	for i := 0; i < len(puzzleInput); i += 2 {

		if i%2 == 0 {
			file_num = helper.ExtractInt(string(puzzleInput[i]))
		} else {
			free_space = helper.ExtractInt(string(puzzleInput[i]))
		}
	}

}

type File struct {
	ID            int
	NumberOfFiles int
	FreeSpace     int
}
