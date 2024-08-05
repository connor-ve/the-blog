package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/components/nav.tmpl.html",
		"../../ui/components/footer.tmpl.html",
		"../../ui/components/info-box.tmpl.html",
		"../../ui/components/skills.tmpl.html",
		"../../ui/components/articles.tmpl.html",
		"../../ui/components/highlights.tmpl.html",
		"../../ui/components/links.tmpl.html",
		"../../ui/components/content.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/components/nav.tmpl.html",
		"../../ui/components/highlights.tmpl.html",
		"../../ui/components/info-box.tmpl.html",
		"../../ui/components/skills.tmpl.html",
		"../../ui/components/articles.tmpl.html",
		"../../ui/components/links.tmpl.html",
		"../../ui/components/footer.tmpl.html",
		"../../ui/components/content.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
