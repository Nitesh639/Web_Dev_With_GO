package main

import (
	"os"
	"text/template"
)

func main(){
	tpl := template.Must(template.ParseFiles("slice_index.html"))

	sl := []int{1,2,3,4,5,6,7,8}

	_ = tpl.ExecuteTemplate(os.Stdout,"slice_index.html",sl)
}