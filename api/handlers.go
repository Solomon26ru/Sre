package api

import (
    "encoding/json"
    "net/http"
    "strconv"
    "todo-app/db"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getTasks(w, r)
    case http.MethodPost:
        addTask(w, r)
    case http.MethodDelete:
        deleteTask(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := db.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
    var task db.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.AddTask(task); err != nil {
        if err.Error() == "task already exists" {
            http.Error(w, err.Error(), http.StatusConflict)  // 409 Conflict
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/tasks/"):]
    taskID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    if err := db.DeleteTask(taskID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
