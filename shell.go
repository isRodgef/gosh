package main

import (
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
		runCmd(shell)
		//runCmd(shell)
	}
}
