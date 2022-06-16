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
	tpl := template.Must(template.ParseFiles("index.html"))
	n1 := Name{
		"Nitesh",
		"Kumar",
	}
	n2 := Name{
		"Anmol",
		"Singh",
	}
	n3 := Name{
		"Utkarsh",
		"Singh",
	}
	n4 := Name{
		"Hemant",
		"Kushwaha",
	}
	n5 := Name{
		"Anjli",
		"Singh",
	}
	n6 := Name{
		"Arpit",
		"Khan",
	}
	sl := []Name{n1,n2,n3,n4,n5,n6}
	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",sl)

}