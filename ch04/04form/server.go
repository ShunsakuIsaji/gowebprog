// リスト4.4
package main

import (
	"fmt"
	"net/http"
)

// POSTのx-www-form-urlencodedを処理する
// URLクエリも入る
// multipart/form-dataは処理できない
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

// URLクエリは入らない
func postformprocess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

func multipartprocess(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024) //multipart/form-dataを処理する、バイト数を指定する
	fmt.Fprintln(w, r.MultipartForm)
}

// エンコード方式にかかわらず、keyを指定して値を取り出すこともできる（最初の１個のみ）
func getvaluefromformkey(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("hello")
	fmt.Fprintf(w, "key:hello, value:%s", value)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", getvaluefromformkey)
	server.ListenAndServe()
}
