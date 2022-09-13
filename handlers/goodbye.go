package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Hello Go")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "opa", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "goodbye %s", d)
}
