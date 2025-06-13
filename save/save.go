package save

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SaveFile(FILE_PATH, task string) (string, error) {
	// Open file with write and append mode, create if it doesn't exist
	f, err := os.OpenFile(FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("We got an error opening the file", err)
		return "We got an error", err
	}

	// Ensure task ends with newline
	if !strings.HasSuffix(task, "\n") {
		task = task + "\n"
	}

	_, err = f.WriteString(task)
	if err != nil {
		fmt.Println("We got an error when writing in file", err)
		return "We got an error", err
	}

	// fmt.Println("Data written successfully!")
	defer f.Close()
	return "Data written successfully!", err
}

func AllTask(FILE_PATH string) []string {
	var tasks []string

	// Try to open the file
	f, err := os.Open(FILE_PATH)
	if err != nil {
		// If the file doesn't exist, return an empty slice
		fmt.Println("Could not open file:", err)
		return tasks
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// Remove trailing newline
		line = strings.TrimSuffix(line, "\n")
		tasks = append(tasks, line)
	}

	return tasks
}

func PrintTasks(FILE_PATH string) {
	f, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Print("No task Present Please add the task")
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}
		// Remove trailing newline
		line = strings.TrimSuffix(line, "\n")
		fmt.Println(line)
	}

}

func MarkDone(FILE_PATH, task string) error {
	// Read all tasks first
	tasks := AllTask(FILE_PATH)

	// Find and mark the task as done
	found := false
	for i := range tasks {
		if tasks[i] == task {
			tasks[i] = tasks[i] + " :# Done"
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task not found: %s", task)
	}

	// Write all tasks back to the file (overwrite the file)
	f, err := os.OpenFile(FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("We got an error opening the file", err)
		return err
	}
	defer f.Close()

	// Write all tasks back
	for _, t := range tasks {
		_, err := f.WriteString(t + "\n")
		if err != nil {
			fmt.Println("We got an error writing to file", err)
			return err
		}
	}

	fmt.Println("Task marked as done successfully!")
	return nil
}
