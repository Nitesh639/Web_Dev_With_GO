package main

import (
	"fmt"
	"os"
	"text/template"
)

func main(){
	tpl,err := template.ParseFiles("first.gohtml") 

	tpl, err = template.ParseFiles("index.html","index1.html")

	err = tpl.ExecuteTemplate(os.Stdout , "first.gohtml",nil)
	err = tpl.ExecuteTemplate(os.Stdout , "index.html",nil)
	err = tpl.ExecuteTemplate(os.Stdout , "index1.html",nil)

	fmt.Println(tpl)
	fmt.Println(err)
}