package main

import (
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseFiles("index.html"))
	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",123456789)
}