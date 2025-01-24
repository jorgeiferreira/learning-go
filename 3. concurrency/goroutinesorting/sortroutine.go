package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, isValid := readUserInput()

	if !isValid {
		return
	}

	//print the sorted data
	print("Unsorted data: ")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])

	}
	println()
	println()
	var splitInto int = 4

	//split the data into 4 slices
	splittedData := splitData(data, splitInto)

	ch := sortDataInSubRoutines(splittedData, splitInto)

	waitSortToFinish(ch, splitInto)

	//merge the data
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			data[i*2+j] = splittedData[i][j]
		}
	}

	fmt.Println()

	sort(data)

	//print the sorted data
	println("Finish Sorted array: ")

	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	//wait for the user to press a key
	fmt.Scanln()

}

func sort(data []int) {
	var output = "Unsorted partition: "
	for i := 0; i < len(data); i++ {
		output += strconv.Itoa(data[i]) + " "
	}
	output += "\n"

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	output += "Sorted partition: "
	for i := 0; i < len(data); i++ {
		output += strconv.Itoa(data[i]) + " "
	}
	output += "\n"
	println(output)
}

func waitSortToFinish(ch chan int, splitInto int) {
	for i := 0; i < splitInto; i++ {
		<-ch
	}
}

func sortDataInSubRoutines(splittedData [][]int, splitInto int) chan int {
	ch := make(chan int)

	for i := 0; i < splitInto; i++ {
		go func(data []int) {
			sort(data)
			ch <- 1
		}(splittedData[i])
	}
	return ch
}

func splitData(data []int, splitInto int) [][]int {
	subSliceSize := len(data) / splitInto

	//println("Data size: ", len(data))
	//println("Subslice size: ", subSliceSize)

	splittedData := make([][]int, 0, splitInto)
	var start = 0
	for i := 0; i < splitInto; i++ {
		end := start + subSliceSize
		//println("i: ", i, "Start: ", start, " End: ", end)
		if i == splitInto-1 {
			splittedData = append(splittedData, data[start:])
		} else {
			splittedData = append(splittedData, data[start:end])
		}
		start += subSliceSize
	}

	return splittedData
}

func readUserInput() ([]int, bool) {
	var intValues = make([]int, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sequence of number separted by space: ")
	userInputs, _ := reader.ReadString('\n')

	userInputs = strings.Replace(userInputs, "\n", "", -1)
	var stringValues = strings.Split(userInputs, " ")
	if len(stringValues) < 8 {
		fmt.Println("Minumum of 8 numbers required")
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
