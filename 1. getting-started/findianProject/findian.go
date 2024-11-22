package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	userInput = strings.ToLower(userInput)
	if strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
