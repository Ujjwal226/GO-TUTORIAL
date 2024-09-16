package main

import (
	"fmt"
)

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 23, "sarah": 45}
	fmt.Println(myMap2["adam"])
	//fmt.Println(myMap2["jason"]) //0- default value of uint8

	var age, ok = myMap2["jason"]
	if ok {
		fmt.Printf("the age is%v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	//delete(myMap2,"adam")

	//for loop for mapd and array
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

}
