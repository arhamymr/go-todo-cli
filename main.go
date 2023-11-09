package main

import (
	"fmt"
	"go-todo-cli/db"
	"os"
)

func init() {
	db.ConnectDB()
}

func commandSelector(cmd, input string) {
	fmt.Println(db.DB)
	switch cmd {
	case "table":
		fmt.Fprintf(os.Stdout, "Table Created: %s \n", input)
	case "insert":
		fmt.Fprintf(os.Stdout, "Data inserted: %s \n", input)
	default:
		fmt.Println("Command not found")
	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s create message \n", os.Args[0])
		os.Exit(1)
	}

	commandSelector(os.Args[1], os.Args[2])
	// Define a flag named "name" with a default value of "Guest"

	// fmt.Println(db.DB, "inside main")

	// db.CreateTable("newTable", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100),email VARCHAR(100)")
	// name := flag.String("name", "Guest", "Your name")

	// // Parse the command-line arguments
	// flag.Parse()

	// // Get the value of the "name" flag
	// nameValue := *name

	// // Print a greeting with the provided name
	// fmt.Printf("Hello, %s!\n", nameValue)

}
