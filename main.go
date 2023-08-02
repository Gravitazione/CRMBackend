package main

import (
	"fmt"

	"gorm.io/gorm"

	"backendCRM/db"
)

func main() {
	db.Connect()

	// test first query
	result := db.Context.Where("id = ?", 1).Table("customer")
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("Customer not found.")
		} else {
			fmt.Println("Error querying customer:", result.Error)
		}
		return
	}

	fmt.Printf("Customer : %v\n", result)

	fmt.Printf("test")
	fmt.Printf("main")
	//
}
