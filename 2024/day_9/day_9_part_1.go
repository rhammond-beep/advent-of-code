package day9

import (
	"fmt"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay9Part1() {
	puzzleInput := "2333133121414131402"
	disk := make([]*MemoryBlock, 0)
	n := len(puzzleInput)

	for i := 0; i < n-1; i += 2 {
		file_number := helper.ExtractInt(string(puzzleInput[i]))
		free_space := helper.ExtractInt(string(puzzleInput[i+1]))
		disk = append(disk, &MemoryBlock{ID: i / 2, NumberOfFiles: file_number, FreeSpace: free_space})
	}

	disk = append(disk, &MemoryBlock{ID: len(disk), NumberOfFiles: helper.ExtractInt(string(puzzleInput[n-1]))})

	for _, file := range disk {
		fmt.Println(file.String())
	}

}

type MemoryBlock struct {
	ID            int
	NumberOfFiles int
	FreeSpace     int
}

func (mb *MemoryBlock) String() string {
	return fmt.Sprintf("ID: %v, Number of Files: %v, FreeSpace: %v", mb.ID, mb.NumberOfFiles, mb.FreeSpace)
}

/*
Move file blocks one at a time from the end of the disk to the leftmost free space block (until
there are no remaining gaps between file blocks). For the disk map 12345, the process will look like the following:

0..111....22222
02.111....2222.
022111....222..
0221112...22...
02211122..2....
022111222......

The stopping condition is where there are no more gaps inbetween memory blocks... i.e. the FreeSpace on each block
is 0.
*/
func compress(memoryBlocks []*MemoryBlock) []int {
	compressed_memory_space := make([]int, 0)
	startingCounter := 0

	for i := len(memoryBlocks) - 1; i > 0; i++ {
		mb
	}
}

/*
 */
func (mb *File) calculateCheckSum() {

}
