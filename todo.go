package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	DONE            = "✅"
	NOT_DONE        = "❌"
	DEVMODE         = true
	TODO_FILENAME   = "todo.csv"
)

func printError(msg string, err error) {
	fmt.Println(msg)
	if DEVMODE && err != nil {
		fmt.Println("Error details:", err)
	}
}

func checkIsFileExist() {
	_, err := os.Stat(TODO_FILENAME)
	if os.IsNotExist(err) {
		initFile()
	}
}

func initFile() error {
	file, err := os.Create(TODO_FILENAME)
	if err != nil {
		printError("Unable to open the csv file", err)
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	size := fileInfo.Size()

	if size == 0 {
		records := [][]string{
			{"ID", "DATE", "TODO", "STATUS", "COMPLETE AT"},
			{"1", time.Now().Format(DATETIME_FORMAT), "Sample todo", NOT_DONE, ""},
		}
		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.WriteAll(records)
	}

	return nil
}

func getTodoCsv() (*os.File, error) {
	file, err := os.OpenFile(TODO_FILENAME, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		printError("Unable to open the file", err)
	}

	return file, err
}

func addTodo(todo string) {
	file, err := getTodoCsv()
	defer file.Close()

	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil || len(records) <= 1 {
		printError("Error reading CSV or CSV is empty", err)
		return
	}

	lastRow := len(records) - 1
	lastId, err := strconv.Atoi(records[lastRow][0])
	if err != nil {
		printError("Unable to parse the ID", err)
		return
	}
	id := lastId + 1

	// Add new record
	newRecord := []string{
		strconv.Itoa(id),
		time.Now().Format(DATETIME_FORMAT),
		todo,
		NOT_DONE,
		"",
	}

	// Reset the file pointer before writing
	// 	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(newRecord)
	if err != nil {
		printError("Error writing to CSV", err)
	}
}

func listTodo() {
	table, err := tablewriter.NewCSV(os.Stdout, "todo.csv", true)
	if err != nil {
		printError("Error reading CSV for listing", err)
		return
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

func toggleDoneStatus(id string) {
	file, err := os.OpenFile(TODO_FILENAME, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		printError("Unable to read the file", err)
		return
	}

	index, err := strconv.Atoi(id)
	if err != nil || index <= 0 {
		printError("Invalid ID", nil)
		return
	}

	for i := 0; i < len(records); i++ {
		if records[i][0] == id {
			if records[i][3] == DONE {
				records[i][3] = NOT_DONE
			} else {
				records[i][3] = DONE
			}
		}
	}

	// Reset the file pointer before writing
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		printError("Error writing to CSV", err)
	}
}

func deleteTodo(id string) {
	file, err := os.OpenFile(TODO_FILENAME, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		printError("Unable tleo read the file", err)
	}

	newRecords := [][]string{}
	for i := 0; i < len(records); i++ {
		if records[i][0] == id {
			continue
		}
		newRecords = append(newRecords, records[i])
	}

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	file.Truncate(0)

	writer.WriteAll(newRecords)
}

func updateTodo(id string, todo string) {
	file, err := os.OpenFile(TODO_FILENAME, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		printError("Unable tleo read the file", err)
	}

	for i := 0; i < len(records); i++ {
		if records[i][0] == id {
			records[i][2] = todo
		}
	}

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(records)
}
