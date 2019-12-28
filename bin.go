package main

import(
	"io/ioutil"
	"net/http"
)

const FileExtension = ".txt"

type Bin struct {
	Id string
	Requests []byte
}

func (bin *Bin) save() error {
	filename := bin.Id + FileExtension
	return ioutil.WriteFile(filename, bin.Requests, 0600)
}

func loadBin(id string) (*Bin, error) {
	filename := id + FileExtension
	requests, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Bin{Id: id, Requests: requests}, nil
}

func (bin *Bin) appendRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	bin.Requests = append(bin.Requests[:], body[:]...)
	return nil
}

