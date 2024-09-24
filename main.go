package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random string
func generateRandomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	// Create a large slice of strings
	var data []string

	// Allocate 10 million random strings
	for i := 0; i < 6500000; i++ {
		data = append(data, generateRandomString(100))
		if i%500000 == 0 {
			fmt.Printf("Allocated %d strings\n", i)
		}
	}

	// Keep the program running to monitor memory usage
	fmt.Println("Finished allocating memory. Press Ctrl+C to exit.")
	select {}
}
