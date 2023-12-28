package main

import (
	"apdex-rule-generator/config"
	"os"
	"text/template"
)

func main() {

	tmpFile := "rule-template.yml"

	tmpl, err := template.New(tmpFile).ParseFiles(tmpFile)
	if err != nil {
		panic(err)
	}

	cfg, err := config.ReadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	output, err := os.Create("output.yml")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	err = tmpl.ExecuteTemplate(output, tmpFile, cfg)
	if err != nil {
		panic(err)
	}
}
