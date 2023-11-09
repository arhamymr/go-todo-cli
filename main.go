package main

import (
	"flag"
	"fmt"
	"go-todo-cli/db"
)


func init() {
	db.ConnectDB()
}

func main() {
	// Define a flag named "name" with a default value of "Guest"

	fmt.Println(db.DB, "inside main")

	db.CreateTable("newTable", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100),email VARCHAR(100)")
	name := flag.String("name", "Guest", "Your name")

	// Parse the command-line arguments
	flag.Parse()

	// Get the value of the "name" flag
	nameValue := *name

	// Print a greeting with the provided name
	fmt.Printf("Hello, %s!\n", nameValue)
}
