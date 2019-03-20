package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {

	reader := bufio.NewReader(os.Stdin)
	//shell := initState("")
	for{
		input, err := reader.ReadString('\n')
		if err != nil{
		
			fmt.Fprintln(os.Stderr,err)
		}
		fmt.Printf("%c",input[1]);
		changeInput(shell,input)
		///fmt.Printf("%s\n",shell.input)
	}
}
