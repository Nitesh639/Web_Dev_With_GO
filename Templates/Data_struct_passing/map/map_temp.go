package main

import (
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseFiles("map_index.html"))

	sqt := map[string]int{"Nitesh":1,"Kush":2,"Hello":4}

	_ = tpl.ExecuteTemplate(os.Stdout,"map_index.html",sqt)
}