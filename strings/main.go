package main

import (
	"fmt"
	"strings"
)

// strings follow utf encoding
func main() {
	var myString string = "resume"
	var index = myString[0]
	fmt.Printf("%v", index)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune rune = 'a'
	fmt.Println(myRune)

	var strSlice = []string{"a", "b"}
	var catstr = ""
	for i := range strSlice {
		catstr += strSlice[i]
	}
	fmt.Printf("\n%v", catstr)

	//string concatination using string builder

	var newstr = []string{"u", "j", "j", "w", "a", "l"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(newstr[i])
	}
	var newcatStr = strBuilder.String()
	fmt.Printf("\n%v", newcatStr)

}
