package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ages := map[string]int{
		"Mike":   23,
		"Donald": 32,
		"Robert": 26,
		"Chris":  45,
		"Grace":  22,
		"Mercy":  20,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ages)
	if err != nil {
		log.Fatalln("Failed to execute template: ", err)
	}
}
