package main

import (
	"fmt"
	"os"
)

const (
	ADD       = "--add"
	LIST      = "--list"
	TODO_DONE = "--done"
	DELETE    = "--delete"
	UPDATE    = "--update"
)

func main() {
	checkIsFileExist()

	switch os.Args[1] {
	case ADD:
		addTodo(os.Args[2])
	case LIST:
		listTodo()
	case TODO_DONE:
		toggleDoneStatus(os.Args[2])
	case UPDATE:
		updateTodo(os.Args[2], os.Args[3])
	case DELETE:
		deleteTodo(os.Args[2])
	default:
		fmt.Println("Invalid argument")
	}
}
