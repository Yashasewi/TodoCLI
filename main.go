package main

import (
	"fmt"
	"os"
	"todo_list/todo"
)

const FILE_PATH = "./task.txt"
const CONFIG_PATH = "./config.json"

func main() {
	command_line_args := os.Args[1:]
	var tasks []string

	if len(command_line_args) == 0 {
		fmt.Println("No command line arguments provided.")
		return
	}

	switch command_line_args[0] {
	case "--add", "-a":
		tasks = todo.CreateTask(tasks, command_line_args[1])
		fmt.Printf("Task added: %s\n", command_line_args[1])

	case "--delete", "-d":
		var deletedIndex int
		tasks, deletedIndex = todo.DeleteTask(tasks, command_line_args[1])
		if deletedIndex >= 0 {
			fmt.Printf("Your task deleted at index %d\n", deletedIndex)
		} else {
			fmt.Println("Task not found")
		}
	case "--print", "-p":
		todo.PrintTasks(tasks)
	case "--help", "-h", "?":
		todo.PrintHelp()
	default:
		fmt.Print("Please use the --help or -h or ? to view the use cases")
	}

	fmt.Println("Current tasks:", tasks)
}
