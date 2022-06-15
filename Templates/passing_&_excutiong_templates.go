package main

import (
	"fmt"
	"os"
	"text/template"
)

func main(){
	tpl , err := template.ParseFiles("first.gohtml")


	fmt.Println(tpl)
	fmt.Println(err)
	nf, err := os.Create("index1.html")
	defer nf.Close()
	err = tpl.Execute(nf,nil)
}