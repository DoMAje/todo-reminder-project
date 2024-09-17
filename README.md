
# Todo Reminder Project

This is a command-line interface (CLI) application to manage your Todo list with support for deadlines, sorting, and more.

## Features
- Add a new Todo item with a specified title and deadline
- Edit the title of an existing Todo item
- Toggle the completion status of a Todo item
- Delete a Todo item
- List all Todo items
- Sort and filter Todo items by title or deadline
- Support for complex time duration formats such as days, hours, and minutes in one line.
- Automatic updating of the remaining time until the deadline for each Todo item

## Commands

### Add a Todo
```bash
--add "Title" --deadline "Xh:Xm:Xs"
```
Example:
```bash
--add "Get Groceries" --deadline "2h 30m"
```

### List Todos
```bash
--list
```

### Toggle a Todo as Completed or Incomplete
```bash
--toggle [index]
```

### Delete a Todo
```bash
--del [index]
```

### Edit a Todo Title
```bash
--edit [index]:"New Title"
```

### Sort Todos
Sort by title or deadline. Use ascending or descending order:
```bash
--sort=[title/deadline] --ascend=[true/false] --list
```

Example:
```bash
--sort=deadline --ascend=false --list
```

## Sorting & Filtering
The application allows you to sort your Todo items either by their title or deadline and display the sorted list. It also allows sorting in ascending or descending order.

## Example Usage
```bash
go run ./ --add "Go outside" --deadline "2h 30m"
go run ./ --list
go run ./ --toggle 1
go run ./ --edit 1:"Go to park"
go run ./ --del 2
go run ./ --sort=deadline --ascend=true --list
```

## Dependencies
- Go 1.18+
- Table library: `github.com/aquasecurity/table`
