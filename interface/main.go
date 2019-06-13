package main

import (
	"fmt"
)

func main() {
	/*
		We'll create pointers to our two implementations of the Storage interface.
		Both satisfy the Storage interface, but are not of that type. They are of their
		own individual types.
	*/

	ms := &MemoryStorage{m: make(map[string]string)}
	rm := &RemoteStorage{}

	/*
	 They can be used as an argument of type Storage, even though they are of their own
	 concrete types
	*/

	PrintStorageDetails(ms)
	PrintStorageDetails(rm)
	/*
		However, since they satisfy the interface by implementing the necessary methods,
		they can be freely assigned to a variable of the Storage interface type
	*/

	var s Storage

	s = rm
	s = ms

	s.Save("a", "b")

	fmt.Println(s.Get("a"))
}
