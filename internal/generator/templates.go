package generator

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

func processTemplate(templateName, outputPath string, data pluginData) {
	// Read template from embedded FS
	tmplData, err := templateFS.ReadFile("templates/" + templateName)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(templateName).Parse(string(tmplData))
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		panic(err)
	}

}
