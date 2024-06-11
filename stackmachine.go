package main

import "errors"

func StackMachine(commands string)(int, error) {
	if commands== ""{
	return 0, errors.New("empty error")
	}
	return 0, errors.New("")
}

func main() {
	// main is unused - run using 
	// go test ./...
}