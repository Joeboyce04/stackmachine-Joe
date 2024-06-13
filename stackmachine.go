package main

import (
"errors" 
"strings"
)

func StackMachine(commands string)(int, error) {
	stack:=[]int{}
	if commands== ""{
	return 0, errors.New("empty error")
	} 
    commandWords := strings.Split(commands, " ")

	for _, command := range commandWords {
		if !ValidCommand(command){
			return 0, errors.New("Invalid Command")
		} 
            }

	/*	for _, command := range commandWords { 
			switch command {
				case "+":
					 if len(stack) < 2 { 
					return 0, errors.New("Too few elements to add") 
				}
				}
			} */

	/*	if command=="+"{
			if len(stack<2){
				return 0, errors.New("Too few elements too")
			}
		}
*/
	
	
	
	

	
		return 0, errors.New("Valid Command")
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