package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/uklineale/go-randutils"
)

const BinEndpoint = "/bin/"
const RequestEndpoint = "/bin/data/" 
const CreateEndpoint = "/create"
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
	id := r.URL.Path[len(BinEndpoint) + len(RequestEndpoint):]
	bin, err := loadBin(id)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1>", "404 Not found")
	}

	err = bin.appendRequest(r)
	bin.save()
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc(BinEndpoint, loadHandler)
	http.HandleFunc(RequestEndpoint, requestHandler)
	http.HandleFunc(CreateEndpoint, createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
