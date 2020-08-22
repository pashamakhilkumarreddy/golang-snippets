package main

import (
	"fmt"
	"strings"
)

const maxLength int = 100

func main() {
	fmt.Println("Stack")
	for true {
		printMessage("Please enter the operation you want to perform on stack", "info")
		printMessage(`
			1. Push
			2. Pop
			3. Print Stack
			4. Peek
			5. Check is stack full
			6. Check is stack empty
		`, "info")
		stack := make([]int, 0)
		var choice uint16
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			printMessage(err.Error(), "error")
		}
		switch choice {
		case 1:
			var value int
			printMessage("Please enter a value to push onto the stack", "info")
			_, err := fmt.Scanf("%d", &value)
			if err != nil {
				printMessage(err.Error(), "error")
			}
			pushed := push(&stack, value)
			if pushed == -1 {
				printMessage("Can't push onto a filled up stack", "warning")
			} else {
				printMessage(
					fmt.Sprintf("Successfully pushed %v onto the stack", value),
					"info")
			}
		case 2:
			popped := pop(&stack)
			if popped == -1 {
				printMessage("Can't pop from an empty stack", "warning")
			} else {
				printMessage(
					fmt.Sprintf("Successfully popped %v from the stack", stack[popped]),
					"info")
			}
		case 3:
			fmt.Printf("%v\n", stack)
		case 4:
			fmt.Println(peek(stack))
		case 5:
			fmt.Println(isFull(&stack))
		case 6:
			fmt.Println(isEmpty(&stack))
		default:
			fmt.Println("If you want to exit, please press y(yes) else n(no)")
			var exitScreen string
			_, err := fmt.Scanf("%q", exitScreen)
			if err != nil {
				printMessage(err.Error(), "error")
			}
			exitScreen = strings.TrimSpace(exitScreen)
			exitScreen = strings.ToLower(exitScreen)
			if exitScreen == "y" || exitScreen == "yes" {
				printMessage("Bye", "info")
				break
			} else {
				continue
			}
		}

	}
}

func printMessage(message, messageType string) {
	switch messageType {
	case "error":
		fmt.Println(message)
	case "warning":
		fmt.Println(message)
	default:
		fmt.Println(message)
	}
}

func push(stack *[]int, value int) int {
	if isFull(stack) == true {
		printMessage("Can't push onto a filled stack", "warning")
		return -1
	}
	_ = append(*stack, value)
	return 0
}

func pop(stack *[]int) int {
	if isEmpty(stack) == true {
		fmt.Println("Can't pop from an empty stack", "warning")
		return -1
	}
	
	return len(*stack) - 1
}

func peek(stack []int) interface{} {
	return stack[len(stack)-1]
}

func isFull(stack *[]int) bool {
	return len(*stack) == maxLength
}

func isEmpty(stack *[]int) bool {
	return len(*stack) == maxLength
}
