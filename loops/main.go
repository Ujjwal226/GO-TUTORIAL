package main

import (
	"fmt"
)

func main() {
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i += 1
		fmt.Printf("The value is: %v \n", i)

	}

	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println(j)
	}

}
