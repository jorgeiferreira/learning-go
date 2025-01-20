package main

import (
	"fmt"
)

/*
The race condition occurs when two or more routines are trying to write to the same variable at the same time.
Since there is no garantee on the order in which the routines are executed, the output of the program is unpredictable.

In the code bellow we have two routines that are being executed concurrently in a for loop.
Both routines are trying to increate the value of the variable data and print the value of data.
The output of this code is unpredictable because there is no way to know the order in which the routines get excuted.
There is no garantee that the output of the first routine will be printed before the output of the next loop iteration.
Everytime that you run the program you will get different outputs.
*/

func main() {
	var data int
	var routin1 = func(value int) {
		data = data + 1
		fmt.Print(fmt.Sprintf("routin1 - iteration No. %d Counter %d\n", value, data))
	}

	var routin2 = func(value int) {
		data = data + 1
		fmt.Print(fmt.Sprintf("routin2 - iteration No. %d Counter %d\n", value, data))
	}

	for i := 1; i <= 6; i++ {
		go routin1(i)
		go routin2(i)
	}

	fmt.Scanln()
}
