package main

import "fmt";


func main() {
	// result := myFunc("Arinze", "Obi")
	// result := add(1, 2, 3, 4, 5)
	defer fmt.Println("This is the last line of the main function: Thanks to defer")
	defer fmt.Println("This is the second to the last line of the main function: Thanks to defer")
	fmt.Println("This is the first line of the main function")
}

func myFunc (argument1 string, argument2 string) string {
	msg := fmt.Sprintf("Hello %s %s", argument1, argument2)
	return msg
}

func add(values ...int) int {
	total := 0;
	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total
}