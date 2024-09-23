package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	memoryHog [][]byte
	mutex     sync.Mutex
)

// Fungsi untuk mengkonsumsi memory sebesar sizeInMB
func consumeMemory(sizeInMB int) {
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < sizeInMB/10; i++ {
		// Menambahkan 10MB data ke slice
		tenMB := make([]byte, 10*1024*1024)
		memoryHog = append(memoryHog, tenMB)
	}
	fmt.Printf("Memory consumed: %d MB\n", len(memoryHog)*10)
}

// Endpoint untuk mulai mengkonsumsi memory
func memoryHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter "size" dari URL (dalam MB)
	sizeParam := r.URL.Query().Get("size")
	if sizeParam == "" {
		http.Error(w, "Please provide size in MB, e.g., /start?size=1024", http.StatusBadRequest)
		return
	}

	// Konversi dari string ke integer
	sizeInMB, err := strconv.Atoi(sizeParam)
	if err != nil || sizeInMB <= 0 {
		http.Error(w, "Invalid size parameter, must be a positive integer", http.StatusBadRequest)
		return
	}

	// Mulai konsumsi memory sebesar sizeInMB
	go consumeMemory(sizeInMB)

	// Berikan respon kepada pengguna
	fmt.Fprintf(w, "Started consuming %d MB of memory\n", sizeInMB)
}

// Endpoint untuk memeriksa status memory yang dikonsumsi
func statusHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	memoryUsed := len(memoryHog) * 10 // Total memory yang dikonsumsi dalam MB
	fmt.Fprintf(w, "Current memory consumption: %d MB\n", memoryUsed)
}

func main() {
	http.HandleFunc("/start", memoryHandler) // Mulai konsumsi memory
	http.HandleFunc("/status", statusHandler) // Cek status penggunaan memory

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
