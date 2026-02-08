// リスト4.2
/*
 1. go run server.go
 2. http://localhost:8080/headers  を表示
*/

package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func getUserAgent(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	ua := h.Get("User-Agent")
	fmt.Fprintln(w, ua)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/useragent", getUserAgent)
	server.ListenAndServe()
}
