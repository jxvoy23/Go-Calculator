package main

import "fmt" // Formatted I/O library

func main() {
	var num1, num2 float64 // variable declaration for floating point
	var operator string    // variable declaration for string

	// Ask the user for first number
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1) // reads the user input and assigns to variable

	// Ask the user for desired operator
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator) // reads the user input and assigns to variable

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2) // reads the user input and assigns to variable

	var result float64 // variable declaration as floating point

	// Perform the operation
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Cannot divide by zero") // throw exception
			return
		}
	default:
		fmt.Println("Invalid operator") // returns if one of the prompted operators is not input
		return
	}

	fmt.Printf("Result: %.2f\n", result) // prints the result to 2 decimal places and a new line
}
