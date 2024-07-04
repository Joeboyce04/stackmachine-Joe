package main


import (
    "errors"
    "strconv"
    "strings"
)


type Stack []int


func NewStack() Stack {
    return Stack{}
}


func (stack *Stack) Pop() (int, error) {
    if len(*stack) == 0 {
        return 0, errors.New("Empty Stack")
    }
    PoppedValue := (*stack)[len(*stack)-1]
    *stack = (*stack)[:len(*stack)-1]
    return PoppedValue, nil
}


func (stack *Stack) Dup() error {
    if len(*stack) == 0 {
        return errors.New("Empty Stack")
    }
    top := (*stack)[len(*stack)-1]
    *stack = append(*stack, top)
    return nil
}

func (stack *Stack) Add() error {
    if len(*stack) < 2 {
        return errors.New("Too Few Elements")
    }
    second, first := (*stack)[len(*stack)-1], (*stack)[len(*stack)-2]
    *stack = (*stack)[:len(*stack)-2]
    if first+second > 50000 {
        return errors.New("Overflow Error")
    }
    *stack = append(*stack, first+second)
    return nil
}


func (stack *Stack) Subtract() error {
    if len(*stack) < 2 {
        return errors.New("Too Few Elements")
    }
    second, first := (*stack)[len(*stack)-1], (*stack)[len(*stack)-2]
    *stack = (*stack)[:len(*stack)-2]
    if second-first < 0 {
        return errors.New("Underflow Error")
    }
    *stack = append(*stack, second-first)
    return nil
}


func (stack *Stack) Multiply() error {
    if len(*stack) < 2 {
        return errors.New("Too Few Elements")
    }
    first, second := (*stack)[len(*stack)-2], (*stack)[len(*stack)-1]
    *stack = (*stack)[:len(*stack)-2]
    if first*second > 5000 {
        return errors.New("Overflow Error")
    }
    *stack = append(*stack, first*second)
    return nil
}


func (stack *Stack) Clear() error {
    *stack = []int{}
    return nil
}


func (stack *Stack) Sum() error {
    if len(*stack) == 0 {
        return errors.New("Empty Stack")
    }
    sum := 0
    for _, num := range *stack {
        if sum+num > 50000 {
            return errors.New("Overflow Error")
        }
        sum += num
    }
    *stack = []int{sum}
    return nil
}





func StackMachine(commands string) (int, error) {
    stack := NewStack()
    if commands == "" {
        return 0, errors.New("empty error")
    }
    commandWords := strings.Split(commands, " ")

    for _, command := range commandWords {

        num, err := strconv.Atoi(command)
        if err == nil {
            if num > 50000 {
                return 0, errors.New("Overflow Error")
            }

            stack = append(stack, num)

        } else if ValidCommand(command) {

            switch command {
                case "POP":
                    _, err := stack.Pop()
                    if err != nil {
                        return 0, err
                    }
                case "DUP":
                    err := stack.Dup()
                    if err != nil {
                        return 0,  err
                    }
                case "+":
                    err := stack.Add()
                    if err != nil {
                        return 0,  err
                    }
                case "-":
                    err := stack.Subtract()
                    if err != nil {
                        return 0,  err
                    }
                case "*":
                    err := stack.Multiply()
                    if err != nil {
                        return 0,  err
                    }
                case "CLEAR":
                    err := stack.Clear()
                    if err != nil {
                        return 0,  err
                    }
                case "SUM":
                    err := stack.Sum()
                    if err != nil {
                        return 0,  err
                    }
                default:
                    return 0, errors.New("Invalid Command")
                }
        } else {
            return 0, errors.New("Invalid Command")
        }
    }
    if len(stack) == 0 {
        return 0, errors.New("Empty Stack")
    }
    return stack[len(stack)-1], nil
}



func ValidCommand(command string) bool {
    validCommands := []string{"POP", "DUP", "+", "-", "*", "SUM", "CLEAR"}
    for _, validCommand := range validCommands {
        if command == validCommand {
            return true
        }
    }
    return false
}

func main() {
	// main is unused - run using 
	// go test ./...
}