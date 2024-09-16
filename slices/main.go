// unlike arrays slices have dynamic sizes and sizes can be changed during programs execution
package main

import "fmt"

func main() {

	//using make func
	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 8) //capacity gets doubled when size of slices increases

	fmt.Println("Slice", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))

	//stock:= make([]string,0,0) -- empty slice

}
