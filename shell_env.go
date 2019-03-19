package main

import (
	"os"
	"fmt"
)

type shell_state struct{
	input			string  
	env 			[]string 
	oldpwd 			string 	
	pwd 		string 
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

func changeInput(state shell_state,input string) {
	state = initState(input);
	fmt.Println(state.input);

	
}


