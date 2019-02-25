package main

import "fmt"
import "os"
import "os/exec"
import "strings"

func   runCmd(cmd string){
	cmd = strings.TrimSuffix(cmd, "\n")
	to_exec := exec.Command(cmd)
	to_exec.Stdout = os.Stdout
        to_exec.Stderr = os.Stderr
	err := to_exec.Run()
	if err != nil{
		fmt.Println("u suck")
	}

}


