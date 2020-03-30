package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../data"
)

//Products handler
type Products struct {
	l *log.Logger
}

//NewProducts creates a Products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	//List of Products
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}
