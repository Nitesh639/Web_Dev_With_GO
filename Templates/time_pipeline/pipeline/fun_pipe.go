package main

import (
	"math"
	"os"
	"text/template"
)

var fm = template.FuncMap{
	"sq" : sq,
	"sq1" : sq1,
}

func main(){
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",9)
}

func sq (num int) float64{
	return math.Sqrt(float64(num))
}
func sq1 (num int) float64{
	return math.Sqrt2
}

