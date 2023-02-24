package rawebclients

import (
	"html/template"
	"net/http"
)

// ToBrowser - Отправляет ответ браузеру
func ToBrowser(w http.ResponseWriter, htmlFile string, data interface{}) {
	var files []string
	files = append(files, htmlFile)
	files = append(files, "./html/base.html")
	html, err := template.ParseFiles(files...)
	check(err)
	err = html.Execute(w, data)
	check(err)
}
