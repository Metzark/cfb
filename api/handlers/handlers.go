package handlers

import (
	"net/http"
	"text/template"
)

func Teams(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("web/html/index.html"))

	tmpl.Execute(w, nil)
}
