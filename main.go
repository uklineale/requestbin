package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/uklineale/go-randutils"
)

const BinEndpoint = "/bin/"
const CreateEndpoint = "/create"


func loadHandler(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(BinEndpoint):]
	bin, err := loadBin(id)
	if err != nil {
		fmt.Fprintf(writer, "<h1>%s</h1>", "404 Not found")
	}

	fmt.Fprintf(writer, "<h1>%s</h1><div>%s</div>", bin.Id, bin.Requests)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	id := randutils.String(16)
	fmt.Println("Making bin with id: ", id)

	bin := &Bin{Id: id, Requests: make([]byte, 100)}
	bin.save()

	http.Redirect(w, r, BinEndpoint + id, http.StatusFound)
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc(BinEndpoint, loadHandler)
	http.HandleFunc(CreateEndpoint, createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
