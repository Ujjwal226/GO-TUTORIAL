package main

type College struct {
	Student_details Student
	Course_details  Course
	Hostel_details  hostel
	Payment_details payment
}

type Student struct {
	Name  string
	RegNo string
	Year  int
}
type Course struct {
	Degree         string
	Specialization string
}
type hostel struct {
	hostel_name string
	hostel_room int
}

type payment struct {
	fees_paid bool
}

func main() {
	var Student1 College
	Student1.Student_details = Student{
		Name: "Zyz",
	}
}
