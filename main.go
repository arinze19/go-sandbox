package main

import (
	"example/projects"
	"fmt"
)

func main() {
	var choice string

	noteTaker := projects.Note{}
	condition := true

	for condition {
		fmt.Println("Welcome to NoteTaker (Beta) What would you like to do today?")
		fmt.Println("1. Add a new note \n2. List all notes \n3. Delete a note \nPress q to exit the application")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			noteTaker.New()
		case "2":
			noteTaker.Read()
		case "3":
			noteTaker.Delete()
		case "q":
			condition = false
			fmt.Println("Thank you for using NoteTaker (Beta), See you next time!!")
		default:
			fmt.Println("Sorry this option is not supported")
		}
	}
}
