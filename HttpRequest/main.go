package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// URLのパターンとハンドラーを渡す
	fmt.Println("Starting Server...")
	http.HandleFunc("/foo", fooHandlerFunc)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// `/foo` へリクエストがきたときに実行されるハンドラー
// http.ResponseWriterにコンテンツを書き込む。
// もちろんHTMLテンプレートを使ったりJSONを返すことも可能

func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/fooページです!!")
}
