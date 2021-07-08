package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
	path := "web/clock/clock.html.tpl"
	template_file_name := "clock.html.tpl"
	t := template.Must(template.ParseFiles(path))

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, template_file_name, time.Now()); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// ハンドラーを登録
	http.HandleFunc("/clock", clockHandler)

	//ディレクトリを指定する
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web/test"))))

	log.Println("Listing...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
