package main;
import (
    "strings"
)

func clear_space(buffer string) []string{
	return strings.Fields(buffer)
}


/// refactor this to use some sort of map
func is_builtin(state shell_state) bool {
	if (state.tokenized[0] == "echo"){
		return true;
	}
	
	if (state.tokenized[0] == "cd"){
		return true;
	}
	
	return false;
}

func runCmd(state shell_state){
	if is_builtin(state){
		///echo(state.input);
	}
}
