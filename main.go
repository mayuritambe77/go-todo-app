package main

import (
	"encoding/json" // For JSON encoding and decoding
	"fmt"           // For formatted I/O
	"io/ioutil"     // For file I/O
	"os"            // For OS-level operations
	"strconv"       // For string conversion
)

// Task represents a to-do task with an ID, name, and completion status
type Task struct {
	ID   int    `json:"id"`   // Unique identifier for the task
	Name string `json:"name"` // Name or description of the task
	Done bool   `json:"done"` // Completion status of the task
}

var tasks []Task            // Slice to hold all tasks
var dataFile = "tasks.json" // File to store tasks

// Load tasks from file
func loadTasks() {
	file, err := ioutil.ReadFile(dataFile) // Read the file
	if err == nil {
		json.Unmarshal(file, &tasks) // Parse JSON data into tasks slice
	}
}

// Save tasks to file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ") // Convert tasks slice to JSON
	ioutil.WriteFile(dataFile, data, 0644)         // Write JSON data to file
}

// Add a new task
func addTask(taskName string) {
	task := Task{ID: len(tasks) + 1, Name: taskName, Done: false} // Create a new task
	tasks = append(tasks, task)                                   // Add task to tasks slice
	saveTasks()                                                   // Save tasks to file
	fmt.Println("✅ Task added successfully!")                     // Print success message
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("📋 No tasks available.") // Print message if no tasks
		return
	}
	for _, task := range tasks {
		status := "❌" // Default status is not done
		if task.Done {
			status = "✅" // Change status if task is done
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Name, status) // Print task details
	}
}

// Mark task as done
func completeTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = true                                  // Mark task as done
			saveTasks()                                           // Save tasks to file
			fmt.Printf("🎯 Task %d marked as complete!\n", taskID) // Print success message
			return
		}
	}
	fmt.Println("❗ Task not found.") // Print error message if task not found
}

// Main function
func main() {
	loadTasks() // Load tasks from file

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
		addTask(os.Args[2]) // Add a new task
	case "list":
		listTasks() // List all tasks
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("❗ Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2]) // Convert task ID to integer
		if err != nil {
			fmt.Println("❗ Invalid task ID.")
			return
		}
		completeTask(id) // Mark task as done
	default:
		fmt.Println("❗ Unknown command.") // Print error message for unknown command
	}
}
