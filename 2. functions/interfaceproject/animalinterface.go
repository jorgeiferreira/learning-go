package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (a *Cow) Eat() {
	fmt.Println("grass")
}
func (a *Cow) Move() {
	fmt.Println("walk")
}
func (a *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (a *Bird) Eat() {
	fmt.Println("worms")
}
func (a *Bird) Move() {
	fmt.Println("fly")
}
func (a *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (a *Snake) Eat() {
	fmt.Println("mice")
}
func (a *Snake) Move() {
	fmt.Println("slither")
}
func (a *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	animalList := make(map[string]Animal)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You can create a new animal or query an existing animal")
	fmt.Println("To create a new animal, type 'newanimal <animal name> <animal type>'")
	fmt.Println("Animal types: cow, bird, snake")
	fmt.Println("To query an existing animal, type 'query <animal name> <information requested>'")
	fmt.Println("Information requested: eat, move, speak")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		inputs := strings.Split(input, " ")
		if len(inputs) != 3 {
			fmt.Println("Invalid input")
			continue
		}
		command := strings.TrimSpace(inputs[0])
		animalName := strings.TrimSpace(inputs[1])

		if command == "newanimal" {
			animalType := strings.TrimSpace(inputs[2])

			switch animalType {
			case "cow":
				animalList[animalName] = new(Cow)
			case "bird":
				animalList[animalName] = new(Bird)
			case "snake":
				animalList[animalName] = new(Snake)
			default:
				fmt.Println("Invalid animal type")
				continue
			}
			fmt.Println("Created it!")

		} else if command == "query" {
			query := strings.TrimSpace(inputs[2])

			animal, ok := animalList[animalName]
			if !ok {
				fmt.Println("Animal not found")
				continue
			}

			switch query {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid information requested")
			}

		} else {
			fmt.Println("Invalid command")
			continue
		}

	}

}
