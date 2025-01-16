package day9

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay9Part2() {
	// puzzleInput := helper.ReadChallengeInputContigious("./day_9/day9_puzzle_input.txt")
	puzzleInput := "2333133121414131402"
	disk := make([]*MemoryBlock, 0)
	n := len(puzzleInput)
	size := 0
	offset := 0

	for i := 0; i < n-1; i += 2 {
		file_number := helper.ExtractInt(string(puzzleInput[i]))
		free_space := helper.ExtractInt(string(puzzleInput[i+1]))
		size += file_number + free_space
		disk = append(disk, &MemoryBlock{ID: i / 2, NumberOfFiles: file_number, FreeSpace: free_space, Offset: offset})
		offset += file_number + free_space
	}
	numberOfFiles := helper.ExtractInt(string(puzzleInput[n-1]))
	disk = append(disk, &MemoryBlock{ID: len(disk), NumberOfFiles: numberOfFiles, Offset: offset})
	size += helper.ExtractInt(string(puzzleInput[n-1]))

	uncompressed := createUncompressedRepresentation(disk, size)
	fmt.Println(uncompressed)
	compressed := compress2(uncompressed, disk)
	fmt.Println(compressed)
	// compressed := []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}

	checkSum := 0

	for i := 0; i < len(compressed)-1; i++ {
		if compressed[i] == -1 {
			continue
		}
		checkSum += (compressed[i] * i)
	}

	fmt.Println(checkSum)
}
