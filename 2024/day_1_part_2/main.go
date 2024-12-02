package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left_list, right_list := ReadChallengeInput("day_1_input.txt")

	sort.Ints(left_list)
	sort.Ints(right_list)

	fmt.Printf("Total Similarity: %v\n", calculateSimilarityScore(left_list, right_list))
}

/*
Accept a filepath input, assuming each record within the file is
delimetered with the newline character '\n', then return the parsed left and
right hand integer list
*/
func ReadChallengeInput(filepath string) (left_list, right_list []int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), " ")
		left_list = append(left_list, extractInt(pair[0]))
		right_list = append(right_list, extractInt(pair[1]))
	}
	return
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
		fmt.Printf("Next instance: %v\n", item_left)
		occurance_map[item_left] += perform_modified_binary_search(item_left, right_list)
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

	if search_space[m] < element {
		perform_modified_binary_search(element, search_space[m:])
	} else if search_space[m] > element {
		perform_modified_binary_search(element, search_space[:m])
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
