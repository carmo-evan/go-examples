package main

import (
	"fmt"
)

// RemoteStorage could be used to store a remote database's connection/information. Left blank for simplicity.
type RemoteStorage struct {
	//imagine you have what it takes to make a remote database connection here
}

// Save gets the value and stores it in a remote database under the provided key
func (rs *RemoteStorage) Save(key, value string) {
	fmt.Println("Imagine this is saving the value to a remote database")
}

// Get retrieves the value from the remote database
func (rs *RemoteStorage) Get(key string) string {
	fmt.Println("Imagine this is retrieving the value from a remote database")
	return ""
}
