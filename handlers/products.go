package handlers

import (
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	//List of Products
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}
