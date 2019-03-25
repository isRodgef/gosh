package main

import (
"os"
"fmt"
)

func    echo(state shell_state){
	for i, s := range state.tokenized {
		if i !=0{
			fmt.Println(s)
		}
	}
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

func setenv(){
}

func env(){
}

func unsetenv(){
}
