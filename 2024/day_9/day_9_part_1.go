package day9

import (
	"fmt"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay9Part1() {

	puzzleInput := "2333133121414131402"
	files := make([]*File, 0)
	n := len(puzzleInput)

	for i := 0; i < n-1; i += 2 {
		file_number := helper.ExtractInt(string(puzzleInput[i]))
		free_space := helper.ExtractInt(string(puzzleInput[i+1]))
		files = append(files, &File{ID: i / 2, NumberOfFiles: file_number, FreeSpace: free_space})
	}

	files = append(files, &File{ID: len(files), NumberOfFiles: helper.ExtractInt(string(puzzleInput[n-1]))})

	for _, file := range files {
		fmt.Println(file.String())
	}

}

type File struct {
	ID            int
	NumberOfFiles int
	FreeSpace     int
}

func (f *File) String() string {
	return fmt.Sprintf("ID: %v, Number of Files: %v, FreeSpace: %v", f.ID, f.NumberOfFiles, f.FreeSpace)
}
