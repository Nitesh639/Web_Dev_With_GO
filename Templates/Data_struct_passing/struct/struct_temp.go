package main

import (
	"os"
	"text/template"
)

type Name struct{
	First string
	Last string
}

func main(){
	tpl := template.Must(template.ParseFiles("struc_var.html"))
	n1 := Name{
		"Nitesh",
		"Kumar",
	}

	_ = tpl.ExecuteTemplate(os.Stdout,"struc_var.html",n1)

}