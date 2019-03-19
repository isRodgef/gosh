package main;
import (
    "strings"
)

func clear_space(buffer string) []string{
	return strings.Fields(buffer)
}

func is_builtin(cmdName string) bool {
	return false;
}

func runCmd(tokenized []string){
	if is_builtin(tokenized[0]){
		echo(tokenized[0]);
	}
}
