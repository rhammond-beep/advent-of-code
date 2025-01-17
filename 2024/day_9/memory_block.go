package day9

import "fmt"

type MemoryBlock struct {
	ID            int
	NumberOfFiles int
	FreeSpace     int
	WriteOffset   int
	ReadOffset    int
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
func compress(memoryBlocks []int) []int {
	i := 0
	j := len(memoryBlocks) - 1

	for {
		if i == j { // no need to keep swapping when we reach the mid point as all elements will have been evaluated.
			break
		}

		if memoryBlocks[i] != -1 { // find a free spot
			i += 1
			continue
		}

		if memoryBlocks[j] == -1 { // find a valid swap candidate
			j -= 1
			continue
		}

		memoryBlocks[i], memoryBlocks[j] = memoryBlocks[j], memoryBlocks[i]
		i += 1
		j -= 1
	}

	return memoryBlocks
}

/*
This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file.
Attempt to move each file * Exactly once in order of decreasing file ID *
starting with the file with the highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the filethe file does not move.
*/
func compress2(denseRepresentation []int, mbs []*MemoryBlock) []int {
	for i := len(mbs) - 1; i > 0; i-- { // for each file
		for j := 0; j < i; j++ { // check and see if the ith file will fit in the jth position
			if mbs[j].FreeSpace >= mbs[i].NumberOfFiles {
				for k := 0; k < mbs[i].NumberOfFiles; k++ {
					w := mbs[j].WriteOffset + k // write index calculation
					r := mbs[i].ReadOffset + k  // read index calculation

					denseRepresentation[w], denseRepresentation[r] = denseRepresentation[r], denseRepresentation[w]
				}
				mbs[j].WriteOffset += mbs[i].NumberOfFiles
				mbs[j].FreeSpace -= mbs[i].NumberOfFiles
				break
			}
		}
	}

	return denseRepresentation
}
