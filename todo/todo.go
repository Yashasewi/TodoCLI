package todo

import (
	"fmt"
	"slices"
	"todo_list/save"
)

// helps delete task
func DeleteTask(tasks []string, task string) ([]string, int) {
	for i := range tasks {
		if tasks[i] == task {
			// Remove the task by creating a new slice without it
			return slices.Delete(tasks, i, i+1), i
		}
	}
	return tasks, -1
}

// helps create task
func CreateTask(tasks []string, task string) []string {
	tasks = append(tasks, task)
	return tasks
}

// helps print tasks
func PrintTasks(FILE_PATH string) {
	if FILE_PATH == "" {
		fmt.Print("No task found")
		return
	}
	save.PrintTasks(FILE_PATH)
}

func MarkDone(FILE_PATH, task string) {
	if FILE_PATH == "" || task == "" {
		fmt.Print("No task found")
		return
	}
	save.MarkDone(FILE_PATH, task)

}

// print the usages and help
func PrintHelp() {
	fmt.Println("Task Manager - Usage Instructions")
	fmt.Println("--------------------------------")
	fmt.Println("Available commands:")
	fmt.Println("  --add, -a <task>       Add a new task")
	fmt.Println("  --delete, -d <task>    Delete a task")
	fmt.Println("  --print, -p            Print all tasks")
	fmt.Println("  --help, -h             Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go --add \"Buy groceries\"")
	fmt.Println("  go run main.go --delete \"Buy groceries\"")
	fmt.Println("  go run main.go --print")
}
