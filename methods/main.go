package main

import "fmt"

func main() {
	var student1 User
	student1.Name = "Ujjwal"
	student1.Email = "aknsj@demd"
	student1.Status = true
	student1.Age = 21
	student1.oneAge = 22
	student1.GetStatus()
	student1.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev" //value is passed by copy
	fmt.Println("Email is: ", u.Email)

}
