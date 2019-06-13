package main

// MemoryStorage implements an in-memory map for temporary storage of values
type MemoryStorage struct {
	m map[string]string
}

// Save assigns the given value to the in-memory map
func (ms *MemoryStorage) Save(key, value string) {
	ms.m[key] = value
}

// Get retrieves the value saved under a certain key
func (ms *MemoryStorage) Get(key string) string {
	return ms.m[key]
}