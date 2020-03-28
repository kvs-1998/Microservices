package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	/* http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s" , d)
	}) */
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()

	http.HandleFunc()

	http.ListenAndServe(":9090", nil)
}
