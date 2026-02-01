package main

/**
chapt.3を写経
**/

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Addr           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	TLSconfig      *tls.Config
	TLSNextProto   map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState      func(net.Conn, http.ConnState)
	ErrorLog       *log.Logger
}

/** ハンドラによってリクエストを処理する
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type SecondHandler struct{}

func (h SecondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the second handler, %s", r.URL.Path[1:])
}

func main() {
	handler := MyHandler{}
	secondhandler := SecondHandler{}
	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/first", &handler)
	http.Handle("/second/", &secondhandler)
	server.ListenAndServe()
}

**/

// 上記と同じことをハンドラ関数で実行する

// ハンドラ関数、ServeHTTPメソッドのような関数を直で書いちゃう
func firsthandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func secondhandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is second handler, %s", r.URL.Path[1:])
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/first", firsthandle)
	http.HandleFunc("/second/", secondhandle)
	server.ListenAndServe()
}
