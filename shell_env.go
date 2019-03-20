package main

import (
	"os"
	"strings"
)

type shell_state struct{
	input			string
	tokenized		[]string 
	env 			[]string 
	oldpwd 			string 	
	pwd 			string
	buitins			map[string]string 
	
}

func initState(input string) shell_state {
	pwd, err := os.Getwd()
	ret := shell_state{}
	if err == nil{
		env  := os.Environ()
		ret.input = input
		ret.env  = env
		ret.pwd = pwd
		ret.oldpwd = ret.pwd
	}
	return ret

}

func changeInput(state *shell_state,input string) {
	*state = initState(strings.TrimSuffix(input, "\n"));	
	state.tokenized = strings.Split(input, " ")
}


