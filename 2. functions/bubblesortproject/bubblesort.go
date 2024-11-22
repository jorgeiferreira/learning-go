package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var intValues = make([]int, 0, 10)

	intValues, isValid := readUserInput(intValues)
	if !isValid {
		return
	}

	bubblsort(intValues)

	fmt.Println("Sorted data: ", intValues)
}

func bubblsort(values []int) {
	var n = len(values)
	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n-i-1; j++ {
			if values[j] > values[j+1] {
				swap(values, j)
			}
		}
	}
}

func swap(values []int, i int) {
	var temp = values[i]
	values[i] = values[i+1]
	values[i+1] = temp
}

func readUserInput(intValues []int) ([]int, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sequence of number separted by space: ")
	userInputs, _ := reader.ReadString('\n')

	userInputs = strings.Replace(userInputs, "\n", "", -1)
	var stringValues = strings.Split(userInputs, " ")
	if len(stringValues) > 10 {
		fmt.Println("Please enter a maximum of 10 numbers")
		return nil, false
	}

	for _, value := range stringValues {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid input")
			return nil, false
		}
		intValues = append(intValues, intValue)
	}
	return intValues, true
}
