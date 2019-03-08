package main

//import (
//	"os"
//)

type state struct{
	input			string  
	env 			[]string 
	oldpwd 			string 	
	currpwd 		string 
}

func initState(input string) state {
	ret := state{}
	return ret

}


func main(){
	initState("ss")
}

