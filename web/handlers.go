package web

import (
    "html/template"
    "log"
    "net/http"
    "todo-app/db"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tasks, err := db.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Tasks:", tasks)  // Отладочный вывод

    tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
    tmpl.Execute(w, tasks)
}
