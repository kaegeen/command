package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	todoList := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To-Do List Application")
	fmt.Println("=======================")
	fmt.Println("Commands:")
	fmt.Println("  add <task> - Add a task to your to-do list")
	fmt.Println("  list       - Show all tasks in your to-do list")
	fmt.Println("  done <id>  - Mark a task as done (removes it)")
	fmt.Println("  exit       - Quit the application")

	for {
		fmt.Print("\nEnter a command: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
				fmt.Println("Error: Task description cannot be empty.")
			} else {
				todoList = append(todoList, parts[1])
				fmt.Printf("Task added: \"%s\"\n", parts[1])
			}

		case "list":
			if len(todoList) == 0 {
				fmt.Println("Your to-do list is empty.")
			} else {
				fmt.Println("To-Do List:")
				for i, task := range todoList {
					fmt.Printf("  %d. %s\n", i+1, task)
				}
			}

		case "done":
			if len(parts) < 2 {
				fmt.Println("Error: Please specify the task number to mark as done.")
			} else {
				var taskNum int
				_, err := fmt.Sscanf(parts[1], "%d", &taskNum)
				if err != nil || taskNum < 1 || taskNum > len(todoList) {
					fmt.Println("Error: Invalid task number.")
				} else {
					fmt.Printf("Task completed: \"%s\"\n", todoList[taskNum-1])
					todoList = append(todoList[:taskNum-1], todoList[taskNum:]...)
				}
			}

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Please use 'add', 'list', 'done', or 'exit'.")
		}
	}
}
