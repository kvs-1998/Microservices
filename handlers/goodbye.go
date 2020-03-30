package handlers

import (
	"log"
	"net/http"
)

//Goodbye creates a simple handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates a new Goodbye handler with
// the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("BYEEEEEE"))
}
