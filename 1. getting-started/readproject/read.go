package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	FName string
	LName string
}

func main() {
	var fileName string
	fmt.Println("Enter the file name: ")

	fmt.Scan(&fileName)

	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading the file")
		return
	}

	var fileContent = strings.Split(string(file), "\n")

	names := make([]Name, 0, len(fileContent))
	for _, line := range fileContent {
		tempNames := strings.Split(line, " ")
		if len(tempNames) != 2 {
			fmt.Println("Invalid name")
			continue
		}
		names = append(names, Name{FName: tempNames[0], LName: tempNames[1]})
	}

	for _, name := range names {
		fmt.Println(name.FName, name.LName)
	}
}
