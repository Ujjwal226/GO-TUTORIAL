package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator = 10
	var denomenator = 0
	var result, remainder, err = intDivision(numerator, denomenator)
	/*if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result is %v", result)
	} else {
		fmt.Printf("The division value is %v and remainder is %v", result, remainder)

	}*/

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v", result)
	default:
		fmt.Printf("The division value is %v and remainder is %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact!")
	case 1, 2:
		fmt.Println("The division was close!")
	}

}

func intDivision(numerator int, denomenator int) (int, int, error) {
	var err error //nil
	if denomenator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err

	}
	var result int = numerator / denomenator
	var remainder int = numerator % denomenator
	return result, remainder, err
}
