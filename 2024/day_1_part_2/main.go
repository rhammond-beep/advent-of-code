package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*

	Calculate a similarity score by adding u[ rach number in the left list
	after multiplying it by the number of times that number appears in the right list.

*
*/
func main() {
	input_file, _ := os.Open("day_1_input.txt")
	similarityScore := 0

	// Annon Cleanup function to close file after use.
	defer func() {
		err := input_file.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		input_data := make([]byte, 1024)
		n, err := input_file.Read(input_data) // Big bang file read (Is that the best way of doing it?)

		if err != nil && err != io.EOF {
			panic(err)
		} else if n == 0 {
			break
		}

		inputString := strings.Trim(string(input_data[:]), "\n") // Get rid of any trailing characters
		sa := strings.Split(inputString, "\n")

		group1 := make([]int, len(sa)/2)
		group2 := make([]int, len(sa)/2)

		for _, pair := range sa {
			pair2Assign := strings.Split(pair, " ")
			group1 = append(group1, extractInt(pair2Assign[0]))
			group2 = append(group2, extractInt(pair2Assign[1]))
		}

		sort.Ints(group1)
		sort.Ints(group2)

		similarityScore += calculateSimilarityScore(group1, group2)
	}

	fmt.Printf("Total Similarity: %v\n", similarityScore)

}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}

func calculateSimilarityScore(left_list, right_list []int) int {
	occurance_map := make(map[int]int)

	score := 0

	for _, item_left := range left_list {

		occurance_map[item_left] += perform_modified_binary_search(item_left, right_list[:])
	}

	for key, value := range occurance_map {
		score += (key * value)
	}

	return score
}

/*
Take an element and check how many times it occurs within a sorted slice.
The trick here is to perform the binary search as normal (halfing the search space)
until the element is found. Because it's a slice with duplicates allowed, we should
check on either side of the element to see if any more are present, if so, work in
the direction found element no longer exists (counting each one) then terminate
*/
func perform_modified_binary_search(element int, search_space []int) int {

	if len(search_space) == 0 {
		return 0 // handle the case where no occurences found
	}

	m := len(search_space) / 2

	current_element := search_space[m]
	if current_element < element {
		perform_modified_binary_search(element, search_space[m:])
	} else if current_element > element {
		perform_modified_binary_search(element, search_space[0:m])
	} else {
		occurences := 1

		for i := 1; search_space[(m-i)] == element; i++ {
			occurences += 1
		}

		for i := 1; search_space[(m+i)] == element; i++ {
			occurences += 1
		}
		return occurences
	}

	return 0 // shouldn't reach this.
}
