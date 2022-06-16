package main

import (
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc" : strings.ToUpper,
	"ct" : firstThree,
}

func firstThree(s string) string  {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type Name struct{
	First string
	Last string
}

type Fun struct{
	Car string
	Number int
}

type Sli struct{
	Into []Name
	Fu []Fun
}

func main(){
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
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


	f1 := Fun{
		"Hundi",
		123,
	}
	f2 := Fun{
		"Fyun",
		123,
	}
	f3 := Fun{
		"Hundi",
		197,
	}
	f4 := Fun{
		"Lipu",
		897,
	}

	val := Sli{
		[]Name{n1,n2,n3,n4},
		[]Fun{f1,f2,f3,f4},
	}
	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",val)

}