package main

import "fmt";


func main() {
	/**
	* Declaring mutliple varialbes 
	*/
	// var (
	// 	foo string = "Hello"
	// 	bar string = "World"
	// )
	// password := "password"
	// const PI = 3.14
	percent := (7.0 / 9) * 100
	sentence := fmt.Sprintf("hex:%x bin:%b", 10 ,10)
	msg := `This is a 
		multiline message that comes from my 
		heart and also a strange place
	`

	fmt.Printf("Percent: %.2f %%\n", percent)
	fmt.Print(sentence, "\n")
	fmt.Print(msg)
}