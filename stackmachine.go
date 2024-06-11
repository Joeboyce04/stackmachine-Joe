package main

import (
"errors" 
"strings"
)

func StackMachine(commands string)(int, error) {
	if commands== ""{
	return 0, errors.New("empty error")
	} 
    commandWords := strings.Split(commands, " ")

    validCommands := []string{"POP", "DUP", "+", "-", "*", "SUM", "CLEAR"}

    for _, command := range commandWords {

        valid := false

        for _, validCommand := range validCommands {

            if command == validCommand {
                valid = true
                
            }
        }

        if !valid {
            return 0, errors.New("Invalid Command")
        }
    }
    return 0, errors.New("Valid Command")
}
func main() {
	// main is unused - run using 
	// go test ./...
}