package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
) // New import"

// The routes() method returns a servemux containing our application routes.
// func (app *application) routes() http.Handler {
// 	mux := http.NewServeMux()

// 	fileServer := http.FileServer(http.Dir("./ui/static/"))
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	mux.HandleFunc("/", app.home)
// 	mux.HandleFunc("/pipe/view", app.pipeView)
// 	mux.HandleFunc("/pipe/create", app.pipeCreate)

// 	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

// 	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
// 	return standard.Then(mux)
// }

func (app *application) routes() http.Handler {
	// Initialize the router.
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Update the pattern for the route for the static files.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// And then create the routes using the appropriate methods, patterns and
	// handlers.
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/pipe/view/:id", app.pipeView)
	router.HandlerFunc(http.MethodGet, "/pipe/create", app.pipeCreate)
	router.HandlerFunc(http.MethodPost, "/pipe/create", app.pipeCreatePost)

	// Create the middleware chain as normal.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Wrap the router with the middleware and return it as normal.
	return standard.Then(router)
}
