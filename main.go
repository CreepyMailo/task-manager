package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	TaskState   string
	CreatedAt   time.Time
	FinishedAt  time.Time
}

var tasks []Task

const storageFile = "tasks.json"

func loadTasks() {
	file, err := os.ReadFile(storageFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(storageFile, data, 0644)
}

func addTask(title string) {
	id := len(tasks)
	task := Task{ID: id, Title: title, TaskState: "not_started", CreatedAt: time.Now()}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Added task:", title)
}

func listTasks() {
	for _, task := range tasks {
		fmt.Printf("[%s] %d: %s (создана %s)\n", task.TaskState, task.ID, task.Title, task.CreatedAt.Format("2006-01-02"))
	}
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].FinishedAt = time.Now()
			tasks[i].TaskState = "done"
			saveTasks()
			fmt.Println("Completed task:", tasks[i].Title)
			return
		}
	}
}

func deleteTask(id int) {
	if id < 0 || id >= len(tasks) {
		fmt.Println("Task not found:", id)
		return
	}
	tasks = append(tasks[:id], tasks[id+1:]...)
	fmt.Println("Task deleted:", tasks[id].Title)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите команду (add/list/done/delete/exit): ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "add":
			var title string
			fmt.Println("Write tittle")
			title, _ = reader.ReadString('\n')
			fmt.Println("Вы ввели:", title)
			addTask(title)
		case "list":
			listTasks()
		case "done":
			var id int
			fmt.Println("Select the number of the completed task")
			fmt.Scanln(&id)
			completeTask(id)
		case "delete":
			var id int
			fmt.Println("Select the task number to delete")
			fmt.Scanln(&id)
			deleteTask(id)
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Command not found")
		}

	}
}
