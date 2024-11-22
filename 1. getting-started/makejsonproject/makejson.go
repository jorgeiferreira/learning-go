package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	values := make(map[string]string)

	fmt.Println("Enter a name: ")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Enter an address: ")
	var address string
	fmt.Scanf("%s", &address)

	values["name"] = name
	values["address"] = address

	data, err := json.Marshal(values)

	if err != nil {
		fmt.Println("Error marshalling the data")
		return
	}

	fmt.Println(string(data))
}
