package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		// command = command[:len(command)-1]
		command = strings.TrimSpace(command)


		if command=="exit 0" {
			// fmt.Println("Exiting...")
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo "){
			text := strings.TrimSpace(command[5:])
			fmt.Println(text)
			continue
		}
		fmt.Printf("%s: command not found\n", command)

	}
	
	
}
