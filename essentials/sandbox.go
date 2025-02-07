package main

import "fmt"

func sandBox() {
	// result := myFunc("Arinze", "Obi")
	// result := add(1, 2, 3, 4, 5)
	// defer fmt.Println("This is the last line of the main function: Thanks to defer")
	// defer fmt.Println("This is the second to the last line of the main function: Thanks to defer")
	age := 35

	getAdultYears(&age)

	fmt.Println("Age:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}

// DeploymentOptions := [3]string{"Development", "Staging", "Production"}

// for index, value := range DeploymentOptions {
// 	fmt.Println(index, value)
// }

// for i := 0; i < len(DeploymentOptions); i++ {
// 	fmt.Println(DeploymentOptions[i])
// }

// func myFunc (argument1 string, argument2 string) string {
// 	msg := fmt.Sprintf("Hello %s %s", argument1, argument2)
// 	return msg
// }

// func add(values ...int) int {
// 	total := 0;
// 	for i := 0; i < len(values); i++ {
// 		total += values[i]
// 	}

// 	return total
// }
