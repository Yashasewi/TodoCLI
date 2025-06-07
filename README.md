# Todo List CLI Application

A simple command-line todo list manager written in Go.

## Features

- Add tasks to your todo list
- Delete tasks from your todo list
- Print all current tasks
- Simple and easy-to-use command-line interface

## Installation

### Prerequisites

- Go 1.24 or higher

### Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/todo_list.git
   cd todo_list
   ```

2. Build the application:

   ```bash
   go build
   ```

## Usage

### Adding a task

```bash
# To add a task to your todo list
go run main.go --add "Your task here"

# Short form
go run main.go -a "Your task here"
```

### Deleting a task

```bash
# Delete a task from your list
go run main.go --delete "Task to delete"

# Short form
go run main.go -d "Task to delete"
```

### Printing all tasks

```bash
# Display all your current tasks
go run main.go --print

# Short form
go run main.go -p
```

### Getting help

```bash
# Display usage information
go run main.go --help

# Short forms
go run main.go -h
go run main.go ?
```

## Project Structure

- `main.go` - Main entry point for the application; handles command-line arguments and flow control
- `todo/todo.go` - Core functionality for managing todos (create, delete, list)
- `save/save.go` - Planned functionality for persistent storage of todos

## How It Works

The application uses a simple command-line interface to manage your todo list:

1. Tasks are stored in memory during program execution
2. Commands are processed through command-line flags
3. The application provides immediate feedback on actions taken

## Future Work

- [ ] Save data in text file for persistence
- [ ] Mark tasks as complete instead of deleting
- [ ] Add due dates to tasks
- [ ] Implement task categories/tags
- [ ] Add priority levels for tasks
- [ ] Create task filtering and sorting options
- [ ] Build data export/import functionality
- [ ] Save data in JSON or YAML format or database(further down the line)
- [ ] Implement unit tests for core functionality
- [ ] Create a TUI (Text User Interface) for better user experience
- [ ] Add workspace support for multiple todo lists
- [ ] Improve error handling and user feedback

## Troubleshooting

If you encounter any issues:

1. Make sure you have Go 1.24+ installed (`go version`)
2. Verify the command syntax (use `--help` flag)
3. Check that file paths are accessible

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
