package main

import (
"os"
"fmt"
)




func    echo(state shell_state){
	fmt.Sprintln(val)
} 

func	cd(state *shell_state) bool{
	value := ""
	if state.tokenized[0] == "-"{
		value = "prev"
		return true
	}else if tokenized[0] == "~"{
		value = "home"
		return true
	}
	if value == ""{
		os.Chdir(dirname)
	}
	return false;
}

