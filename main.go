package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task
var dataFile = "tasks.json"

// Load tasks from file
func loadTasks() {
	file, err := ioutil.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// Save tasks to file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)
}

// Add a new task
func addTask(taskName string) {
	task := Task{ID: len(tasks) + 1, Name: taskName, Done: false}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("✅ Task added successfully!")
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("📋 No tasks available.")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Name, status)
	}
}

// Mark task as done
func completeTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = true
			saveTasks()
			fmt.Printf("🎯 Task %d marked as complete!\n", taskID)
			return
		}
	}
	fmt.Println("❗ Task not found.")
}

// Main function
func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add <task>      - Add a new task")
		fmt.Println("  list             - List all tasks")
		fmt.Println("  done <task_id>   - Mark task as completed")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("❗ Please provide a task description.")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("❗ Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❗ Invalid task ID.")
			return
		}
		completeTask(id)
	default:
		fmt.Println("❗ Unknown command.")
	}
}
