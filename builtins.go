package main

import (
"os"
"fmt"
)




func    echo(val string){
	fmt.Sprintln(val)
} 

func	cd(dirname string) bool{
	value := ""
	if dirname == "-"{
		value = "prev"
	}else if dirname == "~"{
		value = "home"
	}
	if value == ""{
		os.Chdir(dirname)
	}
	return false;
}

