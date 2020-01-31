package main

import (
	"testing"
	"log"
	"os"
)

func TestSaveAndLoad(t *testing.T) {
	expected := Bin{Id: "abc", Requests: []byte("abc")}
	expected.save()
	actual, err := loadBin("abc")
	if err != nil {
		log.Fatal(err)
	}

	if !Equal(expected.Requests, actual.Requests) {
		t.Errorf("Different bins")
	}

	os.Remove("abc.txt")
}

func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
