package main

import (
	"os"
	"text/template"
)
// var tpl *template.Template 
func main(){
	tpl := template.Must(template.ParseFiles("index.html"))

	_ = tpl.ExecuteTemplate(os.Stdout , "index.html","Kushwaha")  // Kushwaha is the value
}