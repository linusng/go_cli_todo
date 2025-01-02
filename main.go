package main

import (
  "encoding/json"
  "fmt"
  "os"
  "strconv"
)

type Task struct {
  Description string
  Completed bool
}

var tasks []Task
const tasksFile = "tasks.json"

func loadTasks() {
  file, err := os.ReadFile(tasksFile)
  if err == nil {
    json.Unmarshal(file, &tasks)
  }
}

func saveTasks() {
  file, _:= json.MarshalIndent(tasks, "", "  ")
  os.WriteFile(tasksFile, file, 0644)
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: todo [command]\nExample commands: add, list")
    return
  }

  loadTasks()
  defer saveTasks()

  switch os.Args[1] {
  case "add":
    if len(os.Args) < 3 {
      fmt.Println("Usage: todo add [task]")
      return
    }
    tasks = append(tasks, Task{Description: os.Args[2]})
    fmt.Println("Task added.")
  case "list":
    for i, task := range tasks {
      status := "[ ]"
      if task.Completed {
        status = "[x]"
      }
      fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
    }
  case "done":
    if len(os.Args) < 3 {
      fmt.Println("Usage: todo one [task number]")
      return
    }
    index, _ := strconv.Atoi(os.Args[2])
    if index > 0 && index <= len(tasks) {
      tasks[index-1].Completed = true
      fmt.Println("Task marked as complete.")
    } else {
      fmt.Println("Invalid task number.")
    }
  case "remove":
    if len(os.Args) < 3 {
      fmt.Println("Usage: todo remove [task number]")
      return
    }
    index, _ := strconv.Atoi(os.Args[2])
    if index > 0 && index <= len(tasks) {
      tasks = append(tasks[:index-1], tasks[index:]...)
      fmt.Println("Task removed.")
    } else {
      fmt.Println("Invalid task number.")
    }
  case "clear":
    tasks = nil
    fmt.Println("All tasks cleared.")
  case "help":
    fmt.Println(`Usage:
    'todo add [task]' - Add new task
    'todo list' - Lists all tasks
    `)
  default:
    fmt.Println("Unknown command. Use 'todo help' for a list of commands.")
  }
}
