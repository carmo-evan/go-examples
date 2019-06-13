package main

import "fmt"

// Storage is a simple storage interface with Save and Get methods.
type Storage interface {
	Save(key, value string)
	Get(string) string
}

// PrintStorageDetails is a simple example of a method that takes an interface as an argument
func PrintStorageDetails(s Storage) {
	fmt.Println(s)
}
