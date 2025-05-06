package main

import (
	"net/http"

	"github.com/justinas/alice"
) // New import"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/pipe/view", app.pipeView)
	mux.HandleFunc("/pipe/create", app.pipeCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standard.Then(mux)
}
