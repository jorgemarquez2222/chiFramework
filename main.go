package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Post struct {
	ID      string `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	slug2 := chi.URLParam(r, "slug2")
	w.Write([]byte(slug + slug2))
}

func respJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func testChannel(w http.ResponseWriter, r *http.Request) {
	var message = make(chan string)

	go func() {
		message <- "Hola"
	}()
	fmt.Println(<-message)
}

func send(ch chan<- string, value string) { // los datos fluyen hacia el canal
	ch <- value
}

func receive(ch <-chan string) { // los datos fluyen fuera del canal
	fmt.Print("Mensaje del canal unidereccional", <-ch)
}

func channelUnidireccional(w http.ResponseWriter, r *http.Request) {
	var ch = make(chan string)
	go receive(ch)
	send(ch, "Holaaaaaaaa")
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	respJson(w, http.StatusOK, post)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/articles/{slug}-{slug2}", getArticle)
	r.Get("/channel", testChannel)
	r.Get("/channelUnidireccional", channelUnidireccional)
	r.Post("/", PostArticle)
	http.ListenAndServe(":3000", r)
}
