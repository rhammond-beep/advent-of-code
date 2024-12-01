package main

import (
	"fmt"
	"os"
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

	// Annon Cleanup function to close file after use.
	defer func() {
		err := input_file.Close()
		if err != nil {
			panic(err)
		}
	}()
}
