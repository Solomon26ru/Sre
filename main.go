package main

import (
    "log"
    "net/http"
    "todo-app/api"
    "todo-app/db"
    "todo-app/web"
)

func main() {
    db.InitDB()

    http.HandleFunc("/tasks", api.TasksHandler)
    http.HandleFunc("/", web.IndexHandler)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
