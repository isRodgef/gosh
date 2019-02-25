package main

import "os"
import "fmt"

func    echo(val string){
	fmt.Sprintln(val)
} 

func	cd(dirname string){
	os.Chdir(dirname)
}

func main(){
	echo ("L");
}
