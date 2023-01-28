package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	canales "github.com/jorgemarquez2222/chiFramework/canales"
	fact "github.com/jorgemarquez2222/chiFramework/fact"
	products "github.com/jorgemarquez2222/chiFramework/products"
	services "github.com/jorgemarquez2222/chiFramework/services"
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

func signalChannel(w http.ResponseWriter, r *http.Request) {
	canales.MetodoSignal()
	var mapa = map[string]string{"nombre": "Jeannelis"}
	mapa["hola"] = "algo"
	respJson(w, http.StatusOK, mapa)
}

func reuqestTest(w http.ResponseWriter, r *http.Request) {
	url := "http://159.65.241.58:3000/products"
	resp := services.Fetch(url)
	var product products.Produt
	err := json.Unmarshal(resp, &product)
	if err != nil {
		fmt.Println("Error_11:", err)
	}
	product.AddAllQauntity(20)
	fmt.Println(len(product.Products))
	product.RemoveProduct(924)
	product.RemoveProduct(933)
	fmt.Println(len(product.Products))
	product.AddElement(1)
	fmt.Println(len(product.Products))
	respJson(w, http.StatusOK, product)
}

func testFact(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	go fact.Fact(1, &wg)
	go fact.Fact(2, &wg2)

	wg.Wait()
	wg2.Wait()

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/articles/{slug}-{slug2}", getArticle)
	r.Get("/channel", testChannel)
	r.Get("/channelUnidireccional", channelUnidireccional)
	r.Get("/metodoSignal", signalChannel)
	r.Post("/", PostArticle)
	r.Get("/reuqestTest", reuqestTest)
	r.Get("/fact", testFact)
	http.ListenAndServe(":3000", r)
}
