package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//create a new slice with a length of 3
	var values = make([]int, 0, 3)
	for {

		//scan for input
		fmt.Println("Enter a number: ")
		var inputValue string
		fmt.Scanf("%s", &inputValue)

		if inputValue == "x" {
			fmt.Println("Exiting...")
			break
		}

		intValue, err := strconv.Atoi(inputValue)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		//append the input value to the slice
		values = append(values, intValue)

		sort.Ints(values)

		fmt.Println("Sorted data: ", values)
	}
}
