package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Task struct {
    ID          int
    Title       string
    Description string
}

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite3", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT
    )`)
    if err != nil {
        log.Fatal(err)
    }
}

func GetTasks() ([]Task, error) {
    rows, err := db.Query("SELECT id, title, description FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func AddTask(task Task) error {
    _, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
    return err
}

func DeleteTask(id int) error {
    _, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
    return err
}
