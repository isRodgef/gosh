package main

import "fmt"
import "os"
import "os/exec"
import "bufio"
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


func main() {
	reader := bufio.NewReader(os.Stdin)
	for{
		input, err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
		}
		runCmd(input)
	}
	
}
