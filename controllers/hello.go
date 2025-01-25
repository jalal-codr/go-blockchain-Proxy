package controllers

import (
	"fmt"
	"net/http"
	"proxy/templates"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "base.html", map[string]string{
		"Title": "Home Page",
	})
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}
