package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"
)

type task struct {
	TaskName   string
	IsComplete bool
	Timestamp  time.Time
	TimeDue    time.Time
}

func initialiseDataFile() {

	var taskData []task
	file, err := os.Create("task_data.dat")

	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(taskData)

	file.Close()

}

func createTask(suppliedTaskName string, arrData []task) {

	var newTask task
	newTask.TaskName = suppliedTaskName
	newTask.IsComplete = false
	newTask.Timestamp = time.Now()

	var durationNum int64
	var durationType string
	fmt.Print("Enter Duration For Task Completion (m, h, d, wk, mn, yr): ")
	_, err := fmt.Scanf("%d %s", &durationNum, &durationType)
	fmt.Println(durationNum, durationType)

	switch durationType {
	case "m":
		newTask.TimeDue = time.Now().Add(time.Minute * time.Duration(durationNum))
	case "h":
		newTask.TimeDue = time.Now().Add(time.Hour * time.Duration(durationNum))
	case "d":
		newTask.TimeDue = time.Now().Add(time.Hour * time.Duration(24*durationNum))
	case "wk":
		newTask.TimeDue = time.Now().Add(time.Hour * time.Duration(168*durationNum))
	case "mn":
		newTask.TimeDue = time.Now().Add(time.Hour * time.Duration(730*durationNum))
	case "yr":
		newTask.TimeDue = time.Now().Add(time.Hour * time.Duration(8760*durationNum))
	default:
		fmt.Println("Invalid Input")
		return
	}

	newSlice := append(arrData, newTask)

	file, err := os.Create("task_data.dat")
	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(newSlice)

	file.Close()

	fmt.Println("Task Added Successfully")

}

func finishTask(suppliedTaskName string, arrData []task) {

	arrayIndex := 0
	found := false

	for !found && arrayIndex < (len(arrData)-1) {
		if arrData[arrayIndex].TaskName == suppliedTaskName {
			arrData[arrayIndex].IsComplete = true
			found = true
		} else {
			arrayIndex += 1
		}
	}

	file, err := os.Create("task_data.dat")
	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(arrData)

	file.Close()

}

func listTask(arrData []task) {

	fmt.Println("Task List:")
	fmt.Println("Deadline | Task Name : Status")
	for _, val := range arrData {
		var completed string
		if val.IsComplete {
			completed = "Complete"
		} else if !val.IsComplete && time.Now().Sub(val.TimeDue) < 0 {
			completed = "Pending"
		} else {
			completed = "Expired"
		}

		fmt.Printf("%v | %s : %s\n", val.TimeDue.Format("2006-01-02 03:04 PM"), val.TaskName, completed)
	}

}

func removeTask(suppliedTaskName string, arrData []task) {

	arrayIndex := 0
	found := false

	for !found && arrayIndex < (len(arrData)-1) {
		if arrData[arrayIndex].TaskName == suppliedTaskName {
			arrData[arrayIndex].IsComplete = true
			found = true
		} else {
			arrayIndex += 1
		}
	}

	newSlice := append(arrData[0:arrayIndex], arrData[arrayIndex+1:len(arrData)]...)
	fmt.Println(newSlice)

	file, err := os.Create("task_data.dat")
	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(newSlice)

	file.Close()

}

func main() {

	var taskArray []task
	inputArgs := os.Args

	file, err := os.Open("task_data.dat")

	// If File Not Found
	if err != nil {
		initialiseDataFile()
		file, err = os.Open("task_data.dat")
		if err != nil {
			log.Fatal(err)
		}
	}

	decoder := gob.NewDecoder(file)
	decoder.Decode(&taskArray)

	file.Close()

	if len(inputArgs) > 1 {

		currentWork := inputArgs[1]

		switch currentWork {
		case "-c":
			createTask(inputArgs[2], taskArray)
		case "-f":
			finishTask(inputArgs[2], taskArray)
		case "-r":
			removeTask(inputArgs[2], taskArray)
		case "-l":
			listTask(taskArray)
		default:
			fmt.Println("Available Commands:\n-l: List All Tasks Added\n-c: Create New Task\n-f: Finish Pending Task\n-r: Remove A Task")
		}

	} else {

		fmt.Println("Invalid User Input")

	}

}
