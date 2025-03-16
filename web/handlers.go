package web

import (
    "html/template"
    "net/http"
    "todo-app/db"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tasks, err := db.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
    tmpl.Execute(w, tasks)
}
