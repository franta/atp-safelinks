package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["url"]
	if !ok || len(keys[0]) < 1 {
		http.NotFound(w, req)
		return
	}
	http.Redirect(w, req, keys[0], http.StatusMovedPermanently)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirect)
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
}