package main

import (
	"github.com/flysnow-org/soha"
	"html/template"
	"log"
	"os"
)

func main() {

	sohaFuncMap := soha.CreateFuncMap()
	const templateText = `{{md5 .}}`

	tmpl, err := template.New("titleTest").Funcs(sohaFuncMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(os.Stdout, 123456)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
