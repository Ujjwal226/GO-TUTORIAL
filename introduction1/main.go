package main

import "fmt"

func main() {
	var intNum int = 32767 //int 8 int 16 int 32 int 64
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 1123.22 //float 32 float 64
	fmt.Println(floatNum)

	var result float32 = floatNum * float32(intNum) //typecasting
	fmt.Println(result)

	var myString string = "Hello"
	fmt.Println(myString)

	fmt.Println(len(myString))

	//rune datatype for representing len of characters
	var myRune rune = 'a'
	fmt.Println(myRune)

	myVar := "text" //another way for representing variable
	fmt.Println(myVar)

	const myConst string = "Hello" //value cannot be changed
	fmt.Println(myConst)

}
