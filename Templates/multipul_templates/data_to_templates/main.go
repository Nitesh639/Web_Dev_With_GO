package main

import (
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	_ = tpl.ExecuteTemplate(os.Stdout,"index.gohtml",42)
}