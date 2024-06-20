package main

import (
"errors" 
"strings"
"strconv"
)

//CHECK YOUR CAPITAL LETTERS !!!!!!!!!!!!!!!!!!!!!

func StackMachine(commands string)(int, error) {
	stack:=[]int{}
	if commands== ""{
	return 0, errors.New("empty error")
	} 
    commandWords := strings.Split(commands, " ")

	for _, command := range commandWords {
		num, err := strconv.Atoi(command) 
		if err== nil{
			stack = append(stack, num)
		

		} else if ValidCommand(command) {

			switch command {
			case "POP":
				if len(stack) == 0 {
					return 0, errors.New("Empty Stack")
				}
				stack = stack[:len(stack)-1]
			case "DUP":
					 if len(stack) == 0 { 
						return 0, errors.New("Empty Stack")
						 } 
			    top := stack[len(stack)-1] 

				 stack = append(stack, top)
			default:
				return 0, errors.New("Valid Command") //For Now
			} 
			} else {
				return 0, errors.New("Invalid Command") //For Now
			}
			if len(stack)==0{
				return 0, errors.New("Empty Stack")
			}
	}
	
	return stack[len(stack)-1], nil
}


	

		



func ValidCommand(command string)bool{
	validCommands := []string{"POP", "DUP", "+", "-", "*", "SUM", "CLEAR"}

    for _, validCommand := range validCommands{
		if command == validCommand{
			return true
		}
}
return false
}
func main() {
	// main is unused - run using 
	// go test ./...
}