package main

import "fmt"

func main() {
	var printValue string = "Hello World"
	printMe(printValue)
	var a int = 10
	var b int = 20
	var result int = intDivision(a, b)
	fmt.Println(result)

	var z, x int = double(a, b)

	fmt.Printf("The value of the integer division is %v with remainder %v", z, x)

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a int, b int) int {
	var result int = a / b
	return result
}

func double(a int, b int) (int, int) {
	var result1 int = a / b
	var result2 int = a % b
	return result1, result2
}
