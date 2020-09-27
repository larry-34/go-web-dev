package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml", "tpl2.gohtml"))
}

func main() {
	names := []string{"Mike", "Donald", "Robert", "Chris", "Grace", "Mercy"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatalln("Failed to execute specified template: ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", names)
	if err != nil {
		log.Fatalln("Failed to execute specified template: ", err)
	}
}
