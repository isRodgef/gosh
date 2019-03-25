package main

import (
"os"
"fmt"
)

func    echo(state shell_state){
	value := state.tokenized[0]
	fmt.Sprintln(value)
} 

func	cd(state *shell_state) bool{
	value :=(*state).tokenized[0]
	if value == "-"{
		value = os.Getenv("OLDPWD")
		return true
	}else if value == "~"{
		value = os.Getenv("HOME")
		os.Chdir(value)
		return true
	}
	if value == ""{
		os.Chdir("")
	}
	return false;
}

