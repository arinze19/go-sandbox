package essentials

import "fmt"

func AccountManager() {
	var choice int

	fmt.Println(
		`Welcome to the Go bank, What would you like to do?
		1. Check balance 
		2. Deposit money
		3. Withdraw money
		4. Exit
	`)
	fmt.Scan(&choice)

	// wonder what happens if i provide a string or something
	fmt.Printf("This is the typeof choice %v", choice)

	switch choice {
	case 1:
		fmt.Println("You chose the option 1")
	case 2:
		fmt.Println("You chose the option 2")
	case 3:
		fmt.Println("You chose the option 3")
	case 4:
		fmt.Println("You chose the option 4")
	default:
		fmt.Print("Your choice is not contained in the list of choices")
	}
}
