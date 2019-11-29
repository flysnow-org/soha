package main

import (
	"fmt"
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
	sliceData2:=[]user{{Name:"张三",Age:12},{Name:"李四",Age:18},{Name:"王武",Age:20}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w,h{"sliceData":sliceData,"sliceData1":sliceData1,"sliceData2":sliceData2})
	})
	fmt.Println("Please visit 127.0.0.1 ")
	http.ListenAndServe(":80", nil)
}

type user struct {
	Name string
	Age int
}
