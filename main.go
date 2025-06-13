package main

import (
	"fmt"
	"os"
	"todo_list/save"
	"todo_list/todo"
)

const FILE_PATH = "./task.txt"
const CONFIG_PATH = "./config.json"

func main() {
	command_line_args := os.Args[1:]
	tasks := save.AllTask(FILE_PATH)

	if len(command_line_args) == 0 {
		fmt.Println("No command line arguments provided.")
		return
	}

	switch command_line_args[0] {
	case "--add", "-a":
		tasks = todo.CreateTask(tasks, command_line_args[1])
		res, err := save.SaveFile(FILE_PATH, command_line_args[1])
		if err != nil {
			fmt.Println(res, err)
		} else {
			fmt.Println(res)
		}
		fmt.Printf("Task added: %s\n", command_line_args[1])

	case "--delete", "-d":
		var deletedIndex int
		tasks, deletedIndex = todo.DeleteTask(tasks, command_line_args[1])
		if deletedIndex >= 0 {
			fmt.Printf("Your task deleted at index %d\n", deletedIndex)
		} else {
			fmt.Println("Task not found")
		}
	case "--done", "-do":
		doneTask := command_line_args[1]
		todo.MarkDone(FILE_PATH, doneTask)
		fmt.Println("This task is now mark as done :", doneTask)

	case "--print", "-p":
		todo.PrintTasks(FILE_PATH)
	case "--help", "-h", "?":
		todo.PrintHelp()
	default:
		fmt.Print("Please use the --help or -h or ? to view the use cases")
	}

}
