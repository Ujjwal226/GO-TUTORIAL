// structs stands for structure is a complex data structure
package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	person_Contact Contact
	Person_Address Address
}

func main() {
	var ujjwal Person
	//fmt.Println("Ujjwal person: ", ujjwal)
	ujjwal.FirstName = "Ujjwal"
	ujjwal.LastName = "Khanna"
	ujjwal.Age = 21
	//fmt.Println("Person is: ", ujjwal)

	//2nd method

	/*person1 := Person{
		FirstName: "Akash",
		LastName:  "Khanna",
		Age:       22,
	}*/
	//fmt.Println("Person is", person1)

	//3rd method
	var person2 = new(Person)
	person2.Age = 21
	person2.FirstName = "ABC"
	person2.LastName = "DEF"
	//fmt.Println("Person is: ", person2)

	var Employee1 Employee
	Employee1.Person_Details = Person{
		FirstName: "Ujjwal",
		LastName:  "khanna ",
		Age:       21,
	}

	Employee1.person_Contact = Contact{
		Email: "contact@gmail.com",
		Phone: "92932938209483094",
	}

	Employee1.Person_Address = Address{
		House: 112,
		Area:  "SDCSC",
		State: "CKHDC",
	}

	fmt.Println("Employee 1 is: ", Employee1)
	fmt.Printf("Name of Employee 1 is %v and address is %v", Employee1.Person_Details.FirstName, Employee1.Person_Address.Area)

}
