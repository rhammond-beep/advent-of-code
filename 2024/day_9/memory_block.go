package day9

import "fmt"

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
This time we want to compress the entire file as opposed to single blocks of the file

The thing that's tricky about this approach, is we need to build the list on the fly...
So I could do a hybrid approach here:
Go through from the starting record, filling in the ID for the size of the file in the block

assuming the block has free space, then check to see if the record will fit based in it's file size

if it doesn't then just move onto the next record from the front, keeping a pointer to the last record we tried and repeat
*/
func compressFiles(memoryBlocks []*MemoryBlock, size int) []int {
	compressed_format := make([]int, size)

	for i := 0; i < size; i++ {
		compressed_format[i] = -1
	}

	i := 0
	j := len(memoryBlocks) - 1
	offset_start := 0
	offset_end := 0

	for {
		if i == j { // when we get to the same record, we're done
			break
		}

		for z := offset_start; z < (memoryBlocks[i].NumberOfFiles + offset_start); z++ {
			compressed_format[z] = memoryBlocks[i].ID
			offset_end += 1
		}
		offset_start = offset_end

		for x := j; x > i; x-- { // try every single file from the highest to the lowest
			if memoryBlocks[x].NumberOfFiles <= memoryBlocks[i].FreeSpace { // load the new file into the free space
				for z := offset_start; z < (memoryBlocks[x].NumberOfFiles + offset_start); z++ {
					compressed_format[z] = memoryBlocks[x].ID
					offset_end += 1
				}
				memoryBlocks[i].FreeSpace -= memoryBlocks[x].NumberOfFiles
				offset_start = offset_end
				j -= 1
			}
		}

		for z := offset_start; z < (memoryBlocks[i].FreeSpace + offset_start); z++ { // add in any free space
			offset_end += 1
		}
		offset_start = offset_end
		i += 1
	}

	return compressed_format
}
