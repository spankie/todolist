package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Task represent a task
type Task struct {
	ID   int64
	Name string
	Done bool
}

// Tasks represents a list of tasks
type Tasks map[int]*Task

const (
	DataFilename = "database.csv"
)

var (
	// AppName   = os.Args[0]
	DataFile  io.ReadWriteCloser
	TaskStore Tasks
)

func init() {
	csvFile, err := os.OpenFile(DataFilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	// assign to global variable so it can be closed in the main function
	DataFile = csvFile
}

// read all the tasks from the datafile and store it in TaskStore memory
func readAllTasksFromFile(file io.Reader) {
	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("error reading csv file: %v\n", err)
	}

	// initialize the in memory store of tasks
	TaskStore = make(Tasks)
	for _, line := range csvLines {
		done, err := strconv.ParseBool(line[2])
		if err != nil {
			fmt.Printf("something went wrong parsing done value: %v\n", err)
			continue
		}
		ID, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("something went wrong parsing ID value: %v\n", err)
			continue
		}
		TaskStore[ID] = &Task{
			ID:   int64(ID),
			Name: line[1],
			Done: done,
		}
	}
}

// read all the tasks that satisfy the isDone argument from the datafile and store it in TaskStore memory
func readAllTasksFromFileByDone(file io.Reader, isdone bool) {
	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("error reading csv file: %v\n", err)
	}

	// initialize the in memory store of tasks
	TaskStore = make(Tasks)
	for _, line := range csvLines {
		done, err := strconv.ParseBool(line[2])
		if err != nil {
			fmt.Printf("something went wrong parsing done value: %v\n", err)
			continue
		}
		ID, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("something went wrong parsing ID value: %v\n", err)
			continue
		}
		if done == isdone {
			TaskStore[ID] = &Task{
				ID:   int64(ID),
				Name: line[1],
				Done: done,
			}
		}
	}
}

// append a task to the tasks data file in the filesystem
func persistTaskToTaskStore(task Task) {
	csvwriter := csv.NewWriter(DataFile)
	t := []string{strconv.FormatInt(task.ID, 10), task.Name, strconv.FormatBool(task.Done)}
	err := csvwriter.Write(t)
	if err != nil {
		log.Fatal(err)
	}

	csvwriter.Flush()
}

// save (persist) the taskStore to the filesystem
func persistTaskStore() {
	if err := os.Truncate(DataFilename, 0); err != nil {
		log.Printf("Failed to delete content of data file: %v", err)
	}

	csvwriter := csv.NewWriter(DataFile)

	for _, task := range TaskStore {
		t := []string{strconv.FormatInt(task.ID, 10), task.Name, strconv.FormatBool(task.Done)}
		err := csvwriter.Write(t)
		if err != nil {
			log.Fatal(err)
		}
	}

	csvwriter.Flush()
}

// update the done field of the task with id `ID` to the `done` argument
func changeTaskDone(done bool, ID int) {
	// read all the data so the one we are looking for can be located and updated
	readAllTasksFromFile(DataFile)
	if task, ok := TaskStore[ID]; ok {
		task.Done = done
		TaskStore[ID] = task
		persistTaskStore()
	} else {
		log.Printf("error: invalid id specified [%d]", ID)
	}
}

// print out all tasks based on isDone argument passed to it
func listTasksByDone(isDone bool) {
	readAllTasksFromFile(DataFile)
	for _, task := range TaskStore {
		if task.Done == isDone {
			fmt.Printf("%d: %s\n", task.ID, task.Name)
		}
	}
}
