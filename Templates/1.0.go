package main

import "fmt"

func main(){
	name := "Nitesh"
	last := "Kushwaha"
	tcp := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>`+ name + `</h1>
	<p>`+last+`</p>
	
	</body>
	</html>`
	fmt.Println(tcp)
}