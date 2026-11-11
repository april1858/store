package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var store = make(map[string]string)

var keyNotFound = errors.New("key not found")

func Put(key, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	v, ok := store[key]
	if !ok {
		return "", keyNotFound
	}
	return v, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}

func helloMuxHendler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloMuxHandler\n"))
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxHendler)
	r.HandleFunc("/v1/{key}", keyValuePutHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
