package generator

import (
	"os"
	"text/template"
)

func processTemplate(templatePath, outputPath string, data pluginData) {
	tmpl, err := template.ParseFiles(templatePath)
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
