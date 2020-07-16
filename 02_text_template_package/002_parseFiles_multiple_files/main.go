package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templates/one.gohtml")
	if err != nil {
		log.Fatalln("An error occured while parsing files: ", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error executing template: ", err)
	}

	tpl, err = tpl.ParseFiles("templates/two.gohtml", "templates/master.gohtml")
	if err != nil {
		log.Fatalln("Error parsing files: ", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error executing template!", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed to execute template one: ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed to execute template two: ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "master.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed to execute master template", nil)
	}
}
