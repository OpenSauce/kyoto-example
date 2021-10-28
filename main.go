package main

import (
	"log"
	"net/http"

	"github.com/yuriizinets/kyoto"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		kyoto.RenderPage(rw, &PageIndex{})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
