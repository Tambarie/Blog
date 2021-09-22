package models

import (
	"html/template"
	"net/http"
)

var templ *template.Template

func Blog(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//files := []string{
	//	"templates/index.html",
	//	"templates/create.html",
	//}

	t, err := template.ParseFiles("templates/index.html")
	if err == nil {
		for _, entry := range Entries {
			t.ExecuteTemplate(rw,entry)
		}

	}

}

