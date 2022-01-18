package main

import (
	"bookstore/controller"
	"html/template"
	"net/http"
)

// IndexHandler 跳转首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", IndexHandler)
	http.HandleFunc("/login", controller.Login)
	http.ListenAndServe(":8080", nil)
}
