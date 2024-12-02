package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var cow = Animal{"cow", "grass", "walk", "moo"}
	var bird = Animal{"bird", "worms", "fly", "peep"}
	var snake = Animal{"snake", "mice", "slither", "hsss"}

	animalList := make(map[string]Animal)
	animalList["cow"] = cow
	animalList["bird"] = bird
	animalList["snake"] = snake
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Imput animal name and the information requested: ")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		inputs := strings.Split(input, " ")
		if len(inputs) != 2 {
			fmt.Println("Invalid input")
			continue
		}
		animalName := strings.TrimSpace(inputs[0])
		animalInfo := strings.TrimSpace(inputs[1])

		animal, ok := animalList[animalName]

		if !ok {
			fmt.Println("Animal not found")
			continue
		}

		fmt.Println("Animal: ", animalName)
		fmt.Println("Information requested: ", animalInfo)

		switch animalInfo {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid information requested")
		}
	}

}
