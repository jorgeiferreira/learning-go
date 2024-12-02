package main

import "fmt"

func main() {
	var acceleration, initialVelocity, initialDisplacement = getUserIntialInput()
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Println("Enter time: ")
	fmt.Scan(&time)

	fmt.Println("Displacement: ", fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}

func getUserIntialInput() (float64, float64, float64) {
	var acceleration, initialVelocity, initialDisplacement float64

	fmt.Println("Enter acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Println("Enter initial velocity: ")
	fmt.Scan(&initialVelocity)

	fmt.Println("Enter initial displacement: ")
	fmt.Scan(&initialDisplacement)

	return acceleration, initialVelocity, initialDisplacement
}
