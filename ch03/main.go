package main

/**
chapt.3を写経
**/

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

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

// チェインさせるlog関数
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Function called -" + name)
		fmt.Fprintf(w, "Logged by log function: %s", name)
		// 元のハンドラ関数を呼び出す
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/first", log(firsthandle))
	http.HandleFunc("/second/", secondhandle)
	server.ListenAndServe()
}
