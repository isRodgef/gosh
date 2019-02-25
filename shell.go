package main

import "fmt"
import "os"
import "bufio"



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
