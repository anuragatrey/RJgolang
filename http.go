package main

import (
	"io"
	"net/http"
)

type myhandler int

func (h myhandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(res, req.URL.RequestURI())
	io.WriteString(res, req.RequestURI)
}

func main() {
	var h myhandler
	http.ListenAndServe(":9000", h)
}
