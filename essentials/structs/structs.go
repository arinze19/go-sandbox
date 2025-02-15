package struct_lesson

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (person *Person) Clear() {
	person.firstName = ""
	person.lastName = ""
	person.age = 0
}

func (person *Person) Greet() {
	message := fmt.Sprintf("Hello, my name is %s %s and I am %d years old\n", person.firstName, person.lastName, person.age)
	fmt.Print(message)
}

// methods attached to a struct should usually be a pointer receiver
func (person *Person) Initiate() {
	// attaching an argument to a struct
	fmt.Print("What is your name?: ")
	fmt.Scanln(&person.firstName)

	fmt.Print("What is your last name?: ")
	fmt.Scanln(&person.lastName)

	fmt.Print("How old are you?: ")
	fmt.Scanln(&person.age)
}

type Admin struct {
	Person
	role string
}

func (admin *Admin) Greet() {
	admin.role = "admin"
	message := fmt.Sprintf("Hello, my name is %s %s and I am %d years old. I am an admin and my role is %s\n", admin.firstName, admin.lastName, admin.age, admin.role)
	fmt.Print(message)
}
