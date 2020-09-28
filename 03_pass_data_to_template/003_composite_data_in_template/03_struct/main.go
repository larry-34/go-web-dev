package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	admin := user{
		Name: "Larry Okongo",
		Age:  23,
	}

	err := tpl.Execute(os.Stdout, admin)
	if err != nil {
		log.Fatalln("Error while executing template: ", err)
	}
}
