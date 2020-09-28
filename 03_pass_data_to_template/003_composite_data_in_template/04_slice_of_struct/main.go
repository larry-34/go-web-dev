package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	users := []user{
		user{
			Name: "Larry Okongo",
			Age:  23,
		},
		user{
			Name: "Lailah Grant",
			Age:  23,
		},
		user{
			Name: "Moses Ali",
			Age:  89,
		},
		user{
			Name: "Yoweri Kaguta Museveni",
			Age:  75,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln("Error executing specified template: ", err)
	}
}
