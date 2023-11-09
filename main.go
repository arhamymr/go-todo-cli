package main

import (
	"flag"
	"fmt"
	"go-todo-cli/db"
)

func main() {
	db.ConnectDB()
	// Define a flag named "name" with a default value of "Guest"
	name := flag.String("name", "Guest", "Your name")

	// Parse the command-line arguments
	flag.Parse()

	// Get the value of the "name" flag
	nameValue := *name

	// Print a greeting with the provided name
	fmt.Printf("Hello, %s!\n", nameValue)
}
