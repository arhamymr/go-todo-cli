package db

import (
	"fmt"
	"log"
)

// CreateTable creates a sample table for demonstration
func CreateTable(tableName, tableData string) {
	fmt.Println(DB)
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS `+ tableName +` (
			`+ tableData +`
		)
	`)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}

// InsertData inserts sample data into the table
func InsertData() {
	_, err := DB.Exec(`
		INSERT INTO your_table (name, email) VALUES
		('John Doe', 'john@example.com'),
		('Jane Doe', 'jane@example.com')
	`)
	if err != nil {
		log.Fatalf("failed to insert data: %v", err)
	}
}

// ReadData retrieves and prints data from the table
func ReadData() {
	rows, err := DB.Query("SELECT id, name, email FROM your_table")
	if err != nil {
		log.Fatalf("failed to query data: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Printf("ID: %d, Name: %s, Email: %s", id, name, email)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("error reading rows: %v", err)
	}
}

// UpdateData updates data in the table
func UpdateData() {
	_, err := DB.Exec("UPDATE your_table SET email = 'updated@example.com' WHERE name = 'John Doe'")
	if err != nil {
		log.Fatalf("failed to update data: %v", err)
	}
}

// DeleteData deletes data from the table
func DeleteData() {
	_, err := DB.Exec("DELETE FROM your_table WHERE name = 'Jane Doe'")
	if err != nil {
		log.Fatalf("failed to delete data: %v", err)
	}
}
