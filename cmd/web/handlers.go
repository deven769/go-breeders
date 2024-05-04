package main

import (
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "hello world")
	app.render(w, "home.page.gohtml", nil)

}
