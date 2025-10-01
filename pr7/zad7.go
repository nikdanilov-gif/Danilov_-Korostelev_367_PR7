package main

import "fmt"

type Task struct {
    ID    int
    Title string
    Done  bool
}

type TaskManager struct {
    tasks []Task
}

func (zadachi *TaskManager) Add(title string) {
    task := Task{ID: len(zadachi.tasks) + 1, Title: title}
    zadachi.tasks = append(zadachi.tasks, task)
    fmt.Println("Добавлено", title)
}

func (zadachi *TaskManager) Delete(id int) {
    for i, task := range zadachi.tasks {
        if task.ID == id {
            zadachi.tasks = append(zadachi.tasks[:i], zadachi.tasks[i+1:]...)
            fmt.Println("Удалено", task.Title)
            return
        }
    }
}

func (zadachi *TaskManager) Complete(id int) {
    for i := range zadachi.tasks {
        if zadachi.tasks[i].ID == id {
            zadachi.tasks[i].Done = true
            fmt.Println("Выполнено", zadachi.tasks[i].Title)
            return
        }
    }
}

func (zadachi *TaskManager) Show() {
    fmt.Println("\nСписок:")
    for _, task := range zadachi.tasks {
        status := " "
        if task.Done {
            status = "Выполнено"
        }
        fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
    }
}

func main() {
    var zadachi TaskManager
    
    zadachi.Add("Выучить типы переменных")
    zadachi.Add("Сделать проект")
    zadachi.Add("Вополнить практическую №6")
    
    zadachi.Complete(1)
    zadachi.Delete(2)
    zadachi.Show()
}