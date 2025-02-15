package projects

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// functionalities
// - display all notes
// - write a note
// - delete a note
type Note struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

func (note *Note) New() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is the title of the note? ")
	title, _ := reader.ReadString('\n')
	note.Title = strings.TrimSpace(title)

	fmt.Println("What is the content of the note? ")
	content, _ := reader.ReadString('\n')
	note.Content = strings.TrimSpace(content)

	note.Date = time.Now()

	// if err != nil {
	// what the difference between using log.Paniclc and using just panic
	// panic is a built-in function that stops the ordinary flow of control and begins panicking.
	// log.Panicln is a function in the log package that prints the arguments to the standard logger and then calls panic.
	// 	log.Panicln("Error occurered while opening the file:", err)
	// }

	var notes []Note

	file, err := os.ReadFile("notes.json")

	if err != nil {
		// create the file
		_, err = os.Create("notes.json")

		if err != nil {
			log.Panicln("Error occured while creating the file:", err)
		}

		file, err = os.ReadFile("notes.json")

		if err != nil {
			log.Panicln("Error occured while opening the file:", err)
		}
	}

	json.Unmarshal(file, &notes)

	if err != nil {
		log.Panicln("Error occured while marshalling the note:", err)
	}

	note.Id = len(notes) + 1

	notes = append(notes, *note)

	data, err := json.Marshal(notes)

	if err != nil {
		log.Panicln("Error occured while marshalling the note:", err)
	}

	err = os.WriteFile("notes.json", data, 0644)

	if err != nil {
		log.Panicln("Error occured while writing to the file:", err)
	}

	fmt.Println("Note saved to file")
}

func (note *Note) Read() {
	var notes []Note

	// read from the notes file
	file, err := os.ReadFile("notes.json")

	if err != nil {
		log.Panicln("You have no notes created yet, Try creating new notes and try again")
	}

	// why does Unmarshal require pointers?
	err = json.Unmarshal(file, &notes)

	if err != nil {
		log.Panicln("Error occured while unmarshalling notes", err)
	}

	for index, note := range notes {
		fmt.Printf("%d. %s \n", index, note.Title)
		fmt.Printf("   %s \n", note.Content)
		fmt.Printf("===================================== \n")
	}
}

func (note *Note) Delete() {
	// get id
	// search notes for id
	fmt.Println("What is the id of the note you would like to delete")
	var id string
	fmt.Scan(&id)

	normalizedId, err := strconv.Atoi(id)

	if err != nil {
		log.Panicln("Please provide a valid number to delete note")
	}

	var notes []Note
	var updatedNotes []Note

	file, err := os.ReadFile("notes.json")

	if err != nil {
		log.Panicln("You have not notes created yet, Try creating new notes and try again")
	}

	err = json.Unmarshal(file, &notes)

	if err != nil {
		log.Panicln("Error occured while unmarshaling note", err)
	}

	for _, note := range notes {
		if note.Id != normalizedId {
			updatedNotes = append(updatedNotes, note)
		}
	}

	// write updated notes into notes.json
	data, err := json.Marshal(updatedNotes)

	if err != nil {
		log.Panicln("Seems there was an issue while updating your notes, please try again later")
	}

	err = os.WriteFile("notes.json", data, 0644)

	if err != nil {
		log.Panicln("Seems there was an issue while trying to update your notes, please try again later")
	}

	fmt.Printf("Removed note with id %d successfully \n", normalizedId)
}
