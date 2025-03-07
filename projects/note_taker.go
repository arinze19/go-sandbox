package projects

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") // macOS/Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

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

	note.Id = rand.Intn(1000000)

	notes = append(notes, *note)

	data, err := json.Marshal(notes)

	if err != nil {
		log.Panicln("Error occured while marshalling the note:", err)
	}

	err = os.WriteFile("notes.json", data, 0644)

	if err != nil {
		log.Panicln("Error occured while writing to the file:", err)
	}

	ClearTerminal()

	fmt.Println("Note saved to file 🎉🥳")
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

	fmt.Println("=====================================")
	fmt.Println("Here are your notes 📝")
	fmt.Println("=====================================")

	for index, note := range notes {
		fmt.Printf("%d.  | %s    | %s | %s  \n", index+1, note.Title, note.Content, note.Date.Format("January 02, 2006 15:04:05"))
	}
	fmt.Println("=====================================")
}

func (note *Note) Delete() {
	// get id
	fmt.Println("What is the title of the note you would like to delete")
	reader := bufio.NewReader(os.Stdin) // use bufio for input

	var title string
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

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
		fmt.Println(note.Title, title, note.Title == title)
		if note.Title != title {
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

	ClearTerminal()

	fmt.Printf("Removed note with title %s successfully 🚮 \n", title)
}
