package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func queryValsHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()

	items := vals["name"]
	if len(items) == 0 {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Name: %s", items[0])
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", p)
}

func errorHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Name not provided")

	w.WriteHeader(http.StatusBadRequest)
}

func errorHandlerr2(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Name not provided", http.StatusBadRequest)
}

type PersonError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func jsonErrorHandler(w http.ResponseWriter, r *http.Request) {
	pErr := &PersonError{
		Error:   "Invalid Request",
		Message: "Name not provided ...",
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(pErr)
	if err != nil {
		log.Println(err)
	}
}

func errFunc() error {
	dbErr := errors.New("connection closed")

	return fmt.Errorf("api error: %w", dbErr)
}

func main1() {
	err := errFunc()
	fmt.Println(err)

	uwErr := errors.Unwrap(err)
	fmt.Println(uwErr)
}

func main2() {
	http.HandleFunc("/", errorHandler1)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	resp, err := http.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(string(body))
}
