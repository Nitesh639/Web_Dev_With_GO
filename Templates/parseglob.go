package main

import (
	"fmt"
	"os"
	"text/template"
)

// for abstract all type files

// func main(){
// 	tpl , err := template.ParseGlob("basic_file/*")

// 	err = tpl.ExecuteTemplate(os.Stdout,"first.gohtml",nil)
// 	err = tpl.ExecuteTemplate(os.Stdout,"index.html",nil)
// 	err = tpl.ExecuteTemplate(os.Stdout,"index1.html",nil)
// 	fmt.Println(tpl)
// 	fmt.Println(err)
// }

// For perticular file pattern

func main(){
	tpl , err := template.ParseGlob("basic_file/*.html")

	err = tpl.ExecuteTemplate(os.Stdout,"first.gohtml",nil)
	err = tpl.ExecuteTemplate(os.Stdout,"index.html",nil)
	err = tpl.ExecuteTemplate(os.Stdout,"index1.html",nil)
	fmt.Println(tpl)
	fmt.Println(err)
}