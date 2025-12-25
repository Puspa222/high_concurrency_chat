package main

import "fmt"

func main() {
	fmt.Printf("Starting chat server...\n")
	done := make(chan struct{})
	<-done
	// fmt.Println("Closing chat server...")
}
