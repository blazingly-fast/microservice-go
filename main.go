package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Go")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "opa", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello %s", d)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye go")
	})
	http.ListenAndServe(":9090", nil)
}
