package main

import (
	"os"
	"text/template"
)

func getme() (string, error) {
	return "test", nil
}

func getme2(s string) (string, error) {
	return s, nil
}

func main() {
	funcMap := template.FuncMap{
		"getme":  getme,
		"getme2": getme2,
	}

	tmpl, _ := template.New("sample.tmpl").Funcs(funcMap).ParseFiles("sample.tmpl")

	tmpl.Execute(os.Stdout, nil)
}
