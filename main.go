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

	if rand.Intn(100) < 50 {
		fmt.Println("Application started successfully (%50!)")
		return
	}

	fmt.Println("Application failed to start (%50!)")
	os.Exit(1)
}
