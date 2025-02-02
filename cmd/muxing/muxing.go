package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods(http.MethodGet)

	router.HandleFunc("/name/{param}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["param"]
		fmt.Fprintf(w, "Hello, %s!", param)
	}).Methods(http.MethodGet)

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err == nil {
			fmt.Fprintf(w, "I got message:\n%s", string(reqBody))
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		a, errA := strconv.Atoi(r.Header.Get("a"))
		b, errB := strconv.Atoi(r.Header.Get("b"))
		if errA == nil && errB == nil {
			sum := a + b
			w.Header().Set("a+b", strconv.Itoa(sum))
		}
	}).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
