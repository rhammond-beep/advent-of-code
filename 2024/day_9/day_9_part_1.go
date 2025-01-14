package day9

import (
	"fmt"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay9Part1() {
	// puzzleInput := "2333133121414131402"
	puzzleInput := "12345"
	disk := make([]*MemoryBlock, 0)
	n := len(puzzleInput)
	size := 0

	for i := 0; i < n-1; i += 2 {
		file_number := helper.ExtractInt(string(puzzleInput[i]))
		free_space := helper.ExtractInt(string(puzzleInput[i+1]))
		size += file_number + free_space
		disk = append(disk, &MemoryBlock{ID: i / 2, NumberOfFiles: file_number, FreeSpace: free_space})
	}

	disk = append(disk, &MemoryBlock{ID: len(disk), NumberOfFiles: helper.ExtractInt(string(puzzleInput[n-1]))})
	size += helper.ExtractInt(string(puzzleInput[n-1]))

	uncompressed_format := createUncompressedRepresentation(disk, size)
	fmt.Println(uncompressed_format)

	// for _, file := range disk {
	// 	fmt.Println(file.String())
	// }

}

type MemoryBlock struct {
	ID            int
	NumberOfFiles int
	FreeSpace     int
}

func (mb *MemoryBlock) String() string {
	return fmt.Sprintf("ID: %v, Number of Files: %v, FreeSpace: %v", mb.ID, mb.NumberOfFiles, mb.FreeSpace)
}

func createUncompressedRepresentation(mbs []*MemoryBlock, size int) []int {
	uncompressed_memory_space := make([]int, size)
	first_ptr := 0
	second_ptr := 0

	for _, mb := range mbs {
		for j := first_ptr; j < first_ptr+mb.NumberOfFiles; j++ {
			uncompressed_memory_space[j] = mb.ID
			second_ptr += 1
		}
		first_ptr = second_ptr

		for j := first_ptr; j < first_ptr+mb.FreeSpace; j++ {
			uncompressed_memory_space[j] = -1
			second_ptr += 1
		}
		first_ptr = second_ptr
	}

	return uncompressed_memory_space
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
// func compress(memoryBlocks []*MemoryBlock) []int {
// 	compressed_memory_space := make([]int, 0)
// 	mi := 0
// 	write_location := 0 // offset into the compressed memory space
// 	swap_location :=
//
// 	for i := len(memoryBlocks) - 1; i > 0; i++ {
//
// 		for j := write_location; j < write_location+memoryBlocks[mi].NumberOfFiles; j++ { // write in the earliest file contents
// 			compressed_memory_space[j] = memoryBlocks[mi].ID
// 		}
//
// 	}
//
// 	return compressed_memory_space
// }
//
