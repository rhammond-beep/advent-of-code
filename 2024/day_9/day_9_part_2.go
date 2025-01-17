package day9

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay9Part2() {
	puzzleInput := helper.ReadChallengeInputContigious("./day_9/day9_puzzle_input.txt")
	disk := make([]*MemoryBlock, 0)
	n := len(puzzleInput)
	size := 0
	offset := 0

	for i := 0; i < n-1; i += 2 {
		file_number := helper.ExtractInt(string(puzzleInput[i]))
		free_space := helper.ExtractInt(string(puzzleInput[i+1]))
		size += file_number + free_space
		disk = append(disk,
			&MemoryBlock{ID: i / 2,
				NumberOfFiles: file_number,
				FreeSpace:     free_space,
				WriteOffset:   offset + file_number,
				ReadOffset:    offset})
		offset += file_number + free_space
	}
	numberOfFiles := helper.ExtractInt(string(puzzleInput[n-1]))
	disk = append(disk, &MemoryBlock{ID: len(disk), NumberOfFiles: numberOfFiles, WriteOffset: offset + numberOfFiles, ReadOffset: offset})
	size += helper.ExtractInt(string(puzzleInput[n-1]))

	uncompressed := createUncompressedRepresentation(disk, size)
	compressed := compress2(uncompressed, disk)

	checkSum := 0

	for i := 0; i < len(compressed)-1; i++ {
		if compressed[i] == -1 {
			continue
		}
		checkSum += (compressed[i] * i)
	}

	fmt.Println(checkSum)
}
