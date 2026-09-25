package main

import (
	"fmt"
	"math/rand"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func main() {

	if rand.Intn(100) < 65 {
		fmt.Println("Application started successfully (%65!)")
		return
	}

	fmt.Println("Application failed to start (%35!)")
	os.Exit(1)
}
