package main

import "fmt"
import "os"
import "os/exec"
import "strings"

func   run_exec(cmd string){
	cmd = strings.TrimSuffix(cmd, "\n")
	fmt.Print(cmd)
	to_exec := exec.Command(cmd)
	to_exec.Stdout = os.Stdout
        to_exec.Stderr = os.Stderr
	err := to_exec.Run()
	if err != nil{
		fmt.Println("u suck")
	}

}


