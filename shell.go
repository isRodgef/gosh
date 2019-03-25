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
		cwd ,er  := os.Getwd()
		if er != nil{
			fmt.Print("Cannot get current cwd>>")
		}
		fmt.Print(cwd,">>")
		input, err := reader.ReadString('\n')
		if err != nil{
		
			fmt.Fprintln(os.Stderr,err)
		}
		changeInput(&shell,input)
		runCmd(shell)
		//runCmd(shell)
	   }
}
