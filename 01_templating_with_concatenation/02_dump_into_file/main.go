package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Larry Okongo"

	tpl := `
		<!Doctype html>
		<html lang="en">
			<head>
				<meta charset="utf-8">
			</head>
			<body>
				<h1>My name is ` + name + `</h1>
			</body>
		</html>
	`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating the specified file!", err)
	}

	io.Copy(file, strings.NewReader(tpl))
}
