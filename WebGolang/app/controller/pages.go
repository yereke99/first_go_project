package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)


func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "html", "startStaticPage.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil{
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil{
		http.Error(rw, err.Error(), 400)
		return
	}

}
