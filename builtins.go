package main

import "os"
import "fmt"

func    echo(val string){
	fmt.Sprintln(val)
} 

func	cd(dirname string){
	value := ""
	if dirname == "-"{
		value = "prev"
	}else if dirname == "~"{
		value = "home"
	}
	if value == ""{
		os.Chdir(dirname)
	}
}

