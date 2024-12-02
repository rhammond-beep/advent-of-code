package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*

	Reconcile the lists

	Step 1) Parse the file into two seperate lists using space as the delimeter
	Step 2) Sort them into ascending order (smallest first)
	Step 3) Linearly walk through list, summing differnce between pairs (abs)
	Step 4) Print result

*
*/
func main() {
	input_file, err := os.Open("day_1_input.txt")

	stat, err := input_file.Stat()
	if err != nil {
		fmt.Printf("Cannont get file stats with following error: %v", err)
	}
	input_data := make([]byte, stat.Size()) // Big bang file read (Is that the best way of doing it?)
	n, err := input_file.Read(input_data)

	if err != nil {
		fmt.Printf("File read errored with following message, %v\n bytes read: %v", err, n)
	} else {

		fmt.Println("File read successfully")
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

	sum := 0.0

	for i := 0; i < len(group1); i++ {
		sum += math.Abs(float64(group1[i] - group2[i]))
	}

	fmt.Printf("Total difference: %f\n", sum)

	// Annon Cleanup function to close file after use.
	defer func() {
		err := input_file.Close()
		if err != nil {
			panic(err)
		}
	}()
}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}
