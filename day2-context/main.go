package main

import (
	"Gee/day1-http-base/base3/Gee"
	"fmt"
	"net/http"
)

func main() {
	r := Gee.New()
	//register in router
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "handler[%q]=%q", k, v)
		}
	})
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL:%s", req.URL)
	})
	r.Run(":9999")
}
