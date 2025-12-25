package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("--------------------------------------------------------------------")
	}
	fmt.Printf("Starting chat server...\n")
	// done := make(chan struct{})
	// <-done
	fmt.Println("Closing chat server...")
}
