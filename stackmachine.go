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
			if num> 50000{
				return 0, errors.New("Overflow Error")
			}
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
			    topSpot := stack[len(stack)-1] 

				 stack = append(stack, topSpot)

			case "+":
				if len(stack)<2{
					return 0, errors.New("Too Few Elements")
				}
				second, first := stack[len(stack)-1], stack[len(stack)-2] 
				stack =stack[:len(stack)-2]
				if first+second>50000{
					return 0, errors.New("Overflow Error")
				}
				stack= append(stack, second+first)

			case "-":
				if len(stack) < 2 {
					return 0, errors.New("Too Few Elements")
				}
			    second, first := stack[len(stack)-1], stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				if second-first<0 {
					return 0, errors.New("Underflow Error")
				}
				stack = append(stack, second- first)

			case "*":
				if len(stack) < 2 {
					return 0, errors.New("Too Few Elements")
				}
				first, second := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				if first*second>5000{
					return 0, errors.New("Overflow Error")
				}
				stack = append(stack, first*second)

			case "CLEAR":
				stack = []int{}

			case "SUM":
				if len(stack)==0{
					return 0, errors.New("Empty Stack")
				}
				sum := 0
				for _, num := range stack {
					if sum+num > 50000{
						return 0, errors.New("Overflow Error")
					}
					sum += num
				}
				stack = []int{sum}
			default:
				return 0, errors.New("Invalid Command")
			}
		}else {
			return 0,errors.New("Invalid Command")
		}
		if len(stack) == 0 {
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