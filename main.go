package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Global variable untuk menyimpan data di memori
var memoryHog [][]byte
var mu sync.Mutex // Mutex untuk mencegah kondisi balapan

// Handler untuk mengisi memori
func stressMemoryHandler(w http.ResponseWriter, r *http.Request) {
	querySize := r.URL.Query().Get("size")
	if querySize == "" {
		querySize = "100" // Default alokasi 100MB
	}

	// Konversi ukuran (size) dari string ke integer
	size, err := strconv.Atoi(querySize)
	if err != nil {
		http.Error(w, "Invalid size parameter", http.StatusBadRequest)
		return
	}

	mu.Lock() // Mengunci saat penulisan
	defer mu.Unlock()

	// Alokasi memory berdasarkan ukuran yang diberikan
	for i := 0; i < size; i++ {
		// Alokasikan 1MB (1024 * 1024 byte) di setiap iterasi
		memoryHog = append(memoryHog, make([]byte, 1024*1024))
	}

	fmt.Fprintf(w, "Memory usage increased by %dMB\n", size)
}

// Handler untuk membersihkan memori
func clearMemoryHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // Mengunci saat pembersihan
	defer mu.Unlock()

	// Kosongkan memori
	memoryHog = [][]byte{}

	fmt.Fprintln(w, "Memory cleared")
}

func main() {
	http.HandleFunc("/stress-memory", stressMemoryHandler)
	http.HandleFunc("/clear-memory", clearMemoryHandler)

	fmt.Println("Memory stress service running on port 8080")
	http.ListenAndServe(":8080", nil)
}
