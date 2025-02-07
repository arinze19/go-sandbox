package essentials

import (
	"fmt"
)

func ProfitCalculator() {
	// ASSIGNMENT
	// 1. Profit calculator
	// 2. ask for revenue, expenses and tax rate
	// 3. calculate earnings before tax and earnings after tax
	// 4. calculate the ratio
	// 5. print EBT, EAT and ratio
	var revenue, expenses, taxRate float64
	var earningsBeforeTax, earningsAfterTax, ratio float64

	// get user input
	fmt.Print("What is your revenue for the year?: ")
	fmt.Scan(&revenue)

	fmt.Println("====================================")

	fmt.Print("What are your expenses for the year?: ")
	fmt.Scan(&expenses)

	fmt.Println("====================================")

	fmt.Print("What is the tax rate in your country?: ")
	fmt.Scan(&taxRate)

	// calculate the inputs
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio = earningsAfterTax / earningsBeforeTax

	fmt.Printf("Earnings Before Tax: %.2f \n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f \n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f%% \n", ratio)
}
