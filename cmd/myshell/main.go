package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	for {
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
		if err != nil{
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		continue
		}
		command = command[:len(command)-1]
		if command == "exit"{
			break
		}
		fmt.Println(command + ": command not found")
		fmt.Print("$ ")
	}
	
	
}
