package main

import (
	"apdex-rule-generator/config"
	"embed"
	"flag"
	"os"
	"text/template"
)

//go:embed rule-template.yml
var ruleTemplate embed.FS

func main() {
	// Define flags for the config and output file paths
	configPath := flag.String("config", "config.yml", "path to the configuration file")
	outputPath := flag.String("output", "output.yml", "path to the output file")
	flag.Parse()

	// Access the embedded rule-template.yml
	tmpl, err := template.New("rule-template.yml").ParseFS(ruleTemplate, "rule-template.yml")
	if err != nil {
		panic(err)
	}

	// Read configuration from the specified path
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// Create output file at the specified path
	output, err := os.Create(*outputPath)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// Execute the template with the configuration and write to the output file
	err = tmpl.Execute(output, cfg)
	if err != nil {
		panic(err)
	}
}
