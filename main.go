package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---- Operator Functions ----

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero - Invalid operation")
	}
	return a / b, nil
}

// ---- Main Program ----

func main() {

	reader := bufio.NewReader(os.Stdin) // used to clear buffer

	var number1, number2 float64
	var operator string

	fmt.Println("Welcome To CLI Calculator")
	fmt.Println("Input Only Numbers")

	for { // loop starts

		// First number
		fmt.Print("Enter the first number: ")
		_, err1 := fmt.Scanln(&number1)
		if err1 != nil {
			fmt.Println("Error: First input is NOT a number.")
			fmt.Println("Please try again.\n")
			reader.ReadString('\n') // <<< FIX: clears invalid input
			continue
		}

		// Second number
		fmt.Print("Enter the second number: ")
		_, err2 := fmt.Scanln(&number2)
		if err2 != nil {
			fmt.Println("Error: Second input is NOT a number.")
			fmt.Println("Please try again.\n")
			reader.ReadString('\n') // <<< FIX: clears invalid input
			continue
		}

		// Operator
		fmt.Print("Enter operator ( + - * / ): ")
		fmt.Scanln(&operator)

		// Perform the operation
		switch operator {
		case "+":
			result := add(number1, number2)
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, result)

		case "-":
			result := subtract(number1, number2)
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, result)

		case "*":
			result := multiply(number1, number2)
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, result)

		case "/":
			result, err := divide(number1, number2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("%v %s %v = %v\n", number1, operator, number2, result)
			}

		default:
			fmt.Println("Invalid operator. Try again.")
			continue
		}

		// Ask user if they want to continue
		var again string
		fmt.Print("\nDo you want to calculate again? (yes/no): ")
		fmt.Scanln(&again)

		if again == "no" || again == "No" {
			fmt.Println("Goodbye!\n")
			break
		}

		fmt.Println()
	}
}