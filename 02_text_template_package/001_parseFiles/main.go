package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Failed to parse the provided file: ", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Failed to execute template: ", err)
	}
}
