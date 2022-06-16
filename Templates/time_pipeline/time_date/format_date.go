package main

import (
	"os"
	"text/template"
	"time"
)
var fm = template.FuncMap{
	"timed" : timely,
	"timed1" : timely1,
}
func main(){
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))

	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",time.Now())
}

func timely(s time.Time) string{
	return s.Format("02-01-2006")
}
func timely1(s time.Time) string{
	return s.Format("kichan")
}