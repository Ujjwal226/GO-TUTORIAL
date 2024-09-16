package main

import "fmt"

const LoginToken string = "efnef" //public

func main() {
	var username string = "Ujjwal"
	fmt.Println(username)
	fmt.Printf("Variable id of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable id of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable id of type: %T \n", smallVal)

	var smallFloat float64 = 255.453439328972837293729837
	fmt.Println(smallFloat)
	fmt.Printf("Variable id of type: %T \n", smallFloat)

	//default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of the type: ", anotherVariable)

	//implicit type

	var website = "www.google.com"
	fmt.Println(website)
	//website=3(cannot be changed later)

	//no var style (only be used inside method)

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Println("Variable is of the type: %T \n", LoginToken)

}
