package main

import (
	"fmt"
	"os"
	"os/exec"
)


func   run_exec(state shell_state){
	fmt.Print(state.tokenized[0])
	to_exec := exec.Command(state.tokenized[0])
	to_exec.Stdout = os.Stdout
        to_exec.Stderr = os.Stderr
	err := to_exec.Run()
	if err != nil{
		fmt.Println("u suck")
	}

}


