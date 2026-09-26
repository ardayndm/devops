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

	if rand.Intn(100) < 55 {
		fmt.Println("Application started successfully (%55!)")
		return
	}

	fmt.Println("Application failed to start (%45!)")
	os.Exit(1)
}
