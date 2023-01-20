package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("server start")

	router := mux.NewRouter().StrictSlash(true)

	var handler http.Handler
	var HandlerFunc http.HandlerFunc
	HandlerFunc = testhandle
	handler = HandlerFunc

	router.Methods("GET").Path("/").Handler(handler)
	router.Methods("POST").Path("/").Handler(handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func testhandle(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}
