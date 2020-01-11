package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/uklineale/go-randutils"
)

const BinEndpoint = "/bin/"
const RequestEndpoint = "/bin/data/" 
const CreateEndpoint = "/create/"
const IdLength = 16


func loadHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(BinEndpoint):]
	bin, err := loadBin(id)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1>", "404 Not found")
	}

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", bin.Id, bin.Requests)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	id := randutils.String(IdLength)
	fmt.Println("Making bin with id: ", id)

	bin := &Bin{Id: id, Requests: make([]byte, 100)}
	bin.save()

	http.Redirect(w, r, BinEndpoint + id, http.StatusFound)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(RequestEndpoint):]
	bin, err := loadBin(id)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1>", "404 Not found")
	}

	err = bin.appendRequest(r)
	bin.save()

	fmt.Fprintf(w, "{\"success\": true}")
}

// Hello world
// TODO Make home page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {

	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s", port)

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	
	http.HandleFunc(BinEndpoint, loadHandler)
	http.HandleFunc(RequestEndpoint, requestHandler)
	http.HandleFunc(CreateEndpoint, createHandler)

	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
