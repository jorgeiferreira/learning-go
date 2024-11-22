package main

import "fmt"

func main() {
	fmt.Println("Imput floating number: ")
	var x float64
	fmt.Scanf("%f", &x)
	fmt.Println("Truncated data: ", int(x))
}
