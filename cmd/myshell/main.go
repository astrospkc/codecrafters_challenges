package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	// Wait for user input
	initCommands()
	for {
		fmt.Fprint(os.Stdout, "$ ")
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		
		inputCmd := strings.Split(strings.TrimSpace(in)," ")
		// the command to execute()
		cmd := inputCmd[0]
		args:= inputCmd[1:]

		// the command created
		cmd_operate := exec.Command(cmd, args...)
		output, err:= cmd_operate.CombinedOutput()
		if err!=nil{
			fmt.Printf(os.Stderr, "Error executing command: %v\n", err)
			return
		}

		fmt.Println(string(output))

		// cmdFn, ok := commands[cmd]
		// if !ok{
		// 	notFound(cmd)
		// }else{
		// 	cmdFn(args)
		// }

	}
}
type (
	cmdFnc func([]string)
)
var commands = make(map[string]cmdFnc)
func registerCommand(cmd string, fn cmdFnc){
	commands[cmd] = fn
}

/* 
{"exit" : exit,
"echo" : echo,
"type": type
}
*/

func initCommands(){
	registerCommand("exit",exit)
	registerCommand("echo", echo)
	registerCommand("type", typer)
}

func notFound(cmd string){
	fmt.Printf("%s: command not found\n",cmd)
}

func exit(args []string){
	if len(args)==0{
		os.Exit(1)
	}

	if code,err:=strconv.Atoi(args[0]); err==nil{
		os.Exit(code)
	}
}

func echo(args []string){
	fmt.Println(strings.Join(args," "))
}

func typer(args []string){
	if len(args)==0{
		fmt.Println("")
	}
	
	_,builtin := commands[args[0]]
	if builtin{
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	paths := strings.Split(os.Getenv("PATH"),":")
	for _, path := range paths{
		fp := filepath.Join(path, args[0])
		// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.
		if _, err := os.Stat(fp); err == nil{
			fmt.Println(fp)
			return
		}
	}
	fmt.Printf("%s not found\n", args[0])

}

