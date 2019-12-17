package main

import(
	"fmt"
	"io/ioutil"
	"testing"
	"cmp"
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

func loadBin(id string) *Bin {
	filename := id + FileExtension
	// TODO Error handling
	requests, _ := ioutil.ReadFile(filename)
	return &Bin{Id: id, Requests: requests}
}


func TestSaveAndLoad(t *testing.T) {
	expected := Bin{Id: "abc", Requests: []byte{65, 66, 67}}
	expected.save()
	actual := loadBin("abc")
	if !cmp.Equal(expected, actual) {
		t.Errorf("")
	}
}

