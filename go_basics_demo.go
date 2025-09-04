// Basic Go syntax and features.

package main

import "fmt"

// Simple Hello World fn
func helloWorld() {
	fmt.Println("Hello, World from Go!")
}

// Fn to add two numbers
func add(a int, b int) int {
	return a + b
}

// Fn to show Go's for loop and if statement
func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}

// Example function
func demo() {
	helloWorld() // Prints Hello World
	fmt.Println("Welcome to Go! Here are some basics:")
	fmt.Printf("2 + 3 = %d\n", add(2, 3)) // Add 2 + 3
	printNumbers(5)                       // Print numbers 1-5 with even/odd
}

// Main function to run the demo
func main() {
	// Run with: go run learning_go.go
	demo()
}
