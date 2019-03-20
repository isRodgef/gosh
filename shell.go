package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {

	reader := bufio.NewReader(os.Stdin)
	shell := initState("")
	for{
		input, err := reader.ReadString('\n')
		if err != nil{
		
			fmt.Fprintln(os.Stderr,err)
		}
		changeInput(&shell,input)
		if (shell.tokenized[0] == "echo"){
			fmt.Printf(shell.tokenized[0])
		}
		runCmd(shell)
		//runCmd(shell)
	}
}
