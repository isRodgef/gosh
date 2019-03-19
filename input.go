package main;
import (
    "strings"
)

func clear_space(buffer string) []string{
	return strings.Fields(buffer)
}


/// refactor this to use some sort of map
func is_builtin(cmdName string) bool {
	if (cmdName == "echo"){
		return true;
	}
	
	if (cmdName == "cd"){
		return true;
	}
	
	return false;
}

func runCmd(tokenized []string){
	if is_builtin(tokenized[0]){
		echo(tokenized[0]);
	}
}
