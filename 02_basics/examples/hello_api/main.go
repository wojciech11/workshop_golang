package main

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	langToMsg map[string]string
	Port      string
}

func (app *App) PostHelloMsgHandler(w http.ResponseWriter, r *http.Request) {

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to read body"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *App) GetHelloMsgHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *App) ListHelloMsgHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *App) DeleteHelloMsg(w http.ResponseWriter, r *http.Request) {
}

func main() {
	app := App{
		langToMsg: make(map[string]string),
	}

	r := chi.NewRouter()
	r.Post("/hello_msg", app.PostHelloMsgHandler)
	r.Get("/say_hello", app.GetHelloMsgHandler)
	r.Get("/hello_msg", app.ListHelloMsgHandler)
	r.Delete("/hello_msg", app.DeleteHelloMsg)
	http.ListenAndServe(":3000", r)
}
