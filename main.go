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

	if rand.Intn(100) < 60 {
		fmt.Println("Application started successfully (%60!)")
		return
	}

	fmt.Println("Application failed to start (%40!)")
	os.Exit(1)
}
