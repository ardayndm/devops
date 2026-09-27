package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func main() {

	randVal := rand.Intn(100)
	if randVal <= 70 {
		fmt.Println("Application started successfully (%70! - Random Value:", randVal, ")")
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Server is healthy!")
		})

		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Failed to start server:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("Application failed to start (%30! - Random Value:", randVal, ")")
	os.Exit(1)
}
