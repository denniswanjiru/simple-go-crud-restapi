package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/denniswanjiru/restapi/app/handler"
	"github.com/gorilla/mux"
)

// App consists of mux Router Instance
type App struct {
	Router *mux.Router
}

// Initialize binds app to a NewRouter instance and executes setRouters
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/api/books", a.handleRequest(handler.GetBooks))
	a.Get("/api/books/{id}", a.handleRequest(handler.GetBook))
	a.Post("/api/books", a.handleRequest(handler.CreateBook))
	a.Put("/api/books/{id}", a.handleRequest(handler.UpdateBook))
	a.Delete("/api/books/{id}", a.handleRequest(handler.DeleteBook))
}

// Get maps out all get request to get request handlers
func (a *App) Get(endpoint string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(endpoint, f).Methods("Get")
}

// Post maps out all get request to post request handlers
func (a *App) Post(endpoint string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(endpoint, f).Methods("Post")
}

// Put maps out all get request to put request handlers
func (a *App) Put(endpoint string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(endpoint, f).Methods("Put")
}

// Delete maps out all get request to delete request handlers
func (a *App) Delete(endpoint string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(endpoint, f).Methods("Delete")
}

// RequestHandlerFunction Structure
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(h RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

// Run spins up a basic server
func (a *App) Run(port string) {
	fmt.Printf("server running on port, %s \n", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}
