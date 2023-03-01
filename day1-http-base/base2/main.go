package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(resp, "URL:%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(resp, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(resp, "404 %s NOT FOUND\n")
	}
}
func main() {
	engin := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engin))
}
