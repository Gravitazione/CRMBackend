package db

import "fmt"

// These can be called inside services
func fetchFromDb() {
	fmt.Println("Query DB to fetch data")
}
