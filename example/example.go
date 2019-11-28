package main

import (
	"github.com/flysnow-org/soha"
	"html/template"
	"net/http"
)

type h map[string]interface{}

func main() {
	tmpl,err:=template.New("index.html").Funcs(soha.CreateFuncMap()).ParseFiles("example/index.html")
	tmpl = template.Must(tmpl, err)

	sliceData:=[]string{"a","b","c","d","e","f","g"}
	sliceData1:=[]string{"d","e","f","g","h","i","j","k"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w,h{"sliceData":sliceData,"sliceData1":sliceData1})
	})
	http.ListenAndServe(":80", nil)
}
