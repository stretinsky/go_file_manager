package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

const (
	readAction   = "read"
	createAction = "create"
	deleteAction = "delete"
)

type action string
type filename string

type userInput struct {
	action   action
	filename filename
}

func getUserInput() userInput {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		log.Fatalf("Incorrect number of arguments. Expected %v received %v", 2, len(argsWithoutProg))
	}

	actions := []string{readAction, createAction, deleteAction}
	if !slices.Contains(actions, argsWithoutProg[0]) {
		log.Fatalf("Unknown action")
	}

	var userInput userInput = userInput{
		action:   action(argsWithoutProg[0]),
		filename: filename(argsWithoutProg[1]),
	}

	return userInput
}

func readFile(filename filename) {
	file, err := os.Open(string(filename))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}

func createFile(filename filename) {
	file, err := os.Create(string(filename))

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
}

func deleteFile(filename filename) {
	e := os.Remove(string(filename))
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	userInput := getUserInput()

	var f func(filename filename)

	switch userInput.action {
	case createAction:
		f = createFile
	case readAction:
		f = readFile
	case deleteAction:
		f = deleteFile
	default:
		log.Fatalf("Unknown action")
	}

	f(userInput.filename)
}
