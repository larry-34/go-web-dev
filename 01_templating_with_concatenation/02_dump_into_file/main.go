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
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Larry's Back!</title>
			</head>
			<body>
				<h1>` + name + `</h1>
			</body>
		</html>
	`
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Failed to create file:-", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(tpl))
}
