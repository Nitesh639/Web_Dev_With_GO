package main

import (
	"os"
	"text/template"
)

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
		[]Name{n1,n2,n3,n4,n5,n6},
		[]Fun{f1,f2,f3,f4},
	}
	_ = tpl.ExecuteTemplate(os.Stdout,"index.html",val)

}