package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type Task struct {
	id int
	title string
	description string
	completed bool
}

type TaskList struct {
	tasks []Task 
}

var RedColor = color.New(color.FgRed)
var YellowColor = color.New(color.FgYellow)
var GreenColor = color.New(color.FgGreen)

var boldRed = RedColor.Add(color.Bold)
var boldYellow = YellowColor.Add(color.Bold)
var boldGreen = GreenColor.Add(color.Bold)

func add(taskList *TaskList) {
	var title, description string
	
	fmt.Printf("Enter the Task Title: ")
	fmt.Scanln(&title)
	if title == "" {
		boldRed.Println("Task title cannot be blank !!!")
		return
	}
	
	fmt.Println("Enter the Task Description:")
	fmt.Scanln(&description)
	
	task := Task{
		id: len(taskList.tasks) + 1,
		title: title,
		description: description,
		completed: false,
	}
	taskList.tasks = append(taskList.tasks, task)
	boldGreen.Println("\nTask added successfully...\n")
}

func findTask(id int, tasks []Task) int {
	for index, task := range tasks {
		if (task.id == id) {
			return index
		}
	}
	return -1
}


func delete(taskList *TaskList) {
	var taskNum int
	fmt.Printf("Enter Task Number to be deleted: ")
	fmt.Scanf("%d", &taskNum)

	if (len(taskList.tasks) == 0) {
		boldRed.Println("No Data Found\n")
		return
	}

	taskIndex := findTask(taskNum, taskList.tasks)
	if taskIndex != -1 {
		taskList.tasks = append(taskList.tasks[:taskIndex], taskList.tasks[taskIndex + 1:]...)
		boldGreen.Println(fmt.Sprintf("\nTask %d deleted successfully\n", taskNum))
	} else {
		boldRed.Println("\nCould not find Task", taskNum)
	}
}

func markAsCompleted(taskList *TaskList) {
	var taskNum int
	fmt.Printf("Enter Task Number to mark as completed: ")
	fmt.Scanf("%d", &taskNum)

	if (len(taskList.tasks) == 0) {
		boldRed.Println("No Data Found\n")
		return
	}

	taskIndex := findTask(taskNum, taskList.tasks)
	if taskIndex != -1 {
		taskList.tasks[taskIndex].completed = true
		boldGreen.Println(fmt.Sprintf("\nTask %d marked as completed\n", taskNum))
	} else {
		boldRed.Println("\nCould not find Task ", taskNum)
	}
}


func display(tasks []Task) {
	if len(tasks) == 0 {
		boldRed.Println("No Tasks found. Please add new task")
		return
	}
	for _, task := range tasks {
		status := "Not Completed"
		boldYellow.Println(fmt.Sprintf("------------------------------------ TASK %d ---------------------------------\n", task.id))
		fmt.Println("title: ", task.title)
		fmt.Println("description: ", task.description)
		if task.completed {
			status = "Completed"
		}
		fmt.Println("status: ", status, "\n")
		boldYellow.Println("------------------------------------------------------------------------------\n")

	}
}

func main() {
	taskList := TaskList{}
	var option int
	boldYellow.Println("\nWelcome to the User Management Application\n")
	for {
		display(taskList.tasks)
		fmt.Println("\nSelect an Option:\n")
		boldGreen.Println("1. Add a Task\n")
		boldGreen.Println("2. Delete Task\n")
		boldGreen.Println("3. Mark Task as Completed\n")
		boldGreen.Println("4. Exit\n")

		fmt.Scanln(&option)
		switch option {
		case 1: 
			add(&taskList)
		case 2: 
			delete(&taskList)
		case 3: 
			markAsCompleted(&taskList)
		case 4:
			boldRed.Println("\nBye Bye ... User Management Application closed\n")
			os.Exit(0)
		default:
			boldRed.Println("\nInvalid option. Please enter valid number\n")
		}
	}
	
}