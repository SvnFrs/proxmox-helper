package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Print("Enter the text you want to type: ")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Switching to target window in 5 seconds...")
	time.Sleep(5 * time.Second)

	// using ydotool to type the text
	cmd := exec.Command("sudo", "ydotool", "type", input)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
