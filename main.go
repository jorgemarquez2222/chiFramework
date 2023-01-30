package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	assertionTypes "github.com/jorgemarquez2222/chiFramework/assertionType"
	canales "github.com/jorgemarquez2222/chiFramework/canales"
	fact "github.com/jorgemarquez2222/chiFramework/fact"
	factoryMethods "github.com/jorgemarquez2222/chiFramework/factoryMethod"
	"github.com/jorgemarquez2222/chiFramework/interfaces"
	POO "github.com/jorgemarquez2222/chiFramework/poo"
	products "github.com/jorgemarquez2222/chiFramework/products"
	recoverPanic "github.com/jorgemarquez2222/chiFramework/recoverPanic"
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

func testFactMultiplexado(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			fact.Fact(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func testPOO(w http.ResponseWriter, r *http.Request) {
	var person = POO.New("Jorge", "jorgemarquez2222@gmail.com")
	var persons POO.Persons
	persons.AddPersonParam(person)
	persons.AddPerson()
	persons.AddPerson()
	persons.RemoveAllPersonByName("otra person")
	respJson(w, http.StatusOK, persons.GetPersons())
}

/*
test de mapas

	if _, ok := animals["gorilla"]; !ok {
		animals["gorilla"] = "gorilla"
	}
*/
func testErrors(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	var e = map[string]string{"data:": string(data)}
	respJson(w, http.StatusOK, e)
}

func testFiles(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Ocurrió un error al crear el archivo", err)
		respJson(w, http.StatusOK, err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("este es una escritura de prueba para archivo nuevo"))
	if err != nil {
		fmt.Println("Ocurrió un error al crear el archivo", err)
		respJson(w, http.StatusOK, err)
		return
	}
	resp := map[string]string{
		"response": "El archivo se creó corectamente",
	}
	respJson(w, http.StatusOK, resp)
}

func testRecoverPanic(w http.ResponseWriter, r *http.Request) {
	values := make([]float32, 0, 0)
	values = append(values, recoverPanic.Division(3, 2))
	values = append(values, recoverPanic.Division(4, 2))
	values = append(values, recoverPanic.Division(310, 5))
	values = append(values, recoverPanic.Division(3, 0))
	values = append(values, recoverPanic.Division(3, 2))
	fmt.Print(values)
	result := map[string][]float32{
		"data": values,
	}
	respJson(w, http.StatusOK, result)

}

func testInterfaces(w http.ResponseWriter, r *http.Request) {
	personAnimal := interfaces.New("JorgeP", "persona")
	personAnimal2 := interfaces.New("JorgeA", "animal")
	interfaces.Exec(personAnimal, "Nuevo")
	interfaces.Exec(personAnimal2, "nuevapersona")
	fmt.Println()
	fmt.Printf("peronaAnimal %T, valor: %s", personAnimal, personAnimal)
	fmt.Println()
	fmt.Printf("peronaAnimal %T, valor: %s", personAnimal2, personAnimal2)
	fmt.Println()
	respJson(w, http.StatusOK, personAnimal)
}

func factoryMethod(w http.ResponseWriter, r *http.Request) {
	var numberMehtod int
	fmt.Println("Seleccione metodo de pago")
	fmt.Println("1. Paypal / 2. Cash / 3. TDC")
	n, err := fmt.Scanln(&numberMehtod)
	if err != nil {
		panic("Ocurrió un error al leer el dato")
	}
	if n < 1 || n > 3 {
		panic("Debe ingresar un dato válido")
	}
	method3 := factoryMethods.GetMethod(numberMehtod)
	fmt.Printf("Resputa %T\n", method3)

	respJson(w, http.StatusOK, string([]byte("respuesta backend")))
}

func assertionType(w http.ResponseWriter, r *http.Request) {
	assertionTypes.Exec("Jorge marquez")
	data := map[string]string{
		"data": "Respusta correcta",
	}
	respJson(w, http.StatusOK, data)
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
	r.Get("/fact", testFactMultiplexado)
	r.Get("/testPOO", testPOO)
	r.Get("/testErrors", testErrors)
	r.Get("/testFiles", testFiles)
	r.Get("/testRecoverPanic", testRecoverPanic)
	r.Get("/testInterfaces", testInterfaces)
	r.Get("/factoryMethod", factoryMethod)
	r.Get("/assertionType", assertionType)

	http.ListenAndServe(":3000", r)
}
