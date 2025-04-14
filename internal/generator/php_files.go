package generator

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type pluginData struct {
	PluginName      string
	Description     string
	Version         string
	Author          string
	TextDomain      string
	NamespaceName   string
	PackageName     string
	ClassPrefix     string
	ConstantPrefix  string
	FunctionPostfix string
}

func generatePluginData(projectName string) (*pluginData, error) {
	return &pluginData{
		PluginName:      projectName,
		Description:     "A WordPress plugin",
		Version:         "0.1.0",
		Author:          "Your Name",
		TextDomain:      "TextDomain",
		NamespaceName:   "Namespace\\PackageName",
		PackageName:     "PackageName",
		ClassPrefix:     "ClassPrefix",
		ConstantPrefix:  "ConstantPrefix",
		FunctionPostfix: "function_postfix",
	}, nil
}

// Takes the last part of the projectName after '/' and coverts it to a textdomain
func generateTextDomain(projectName string) (*string, error) {
	parts := strings.Split(projectName, "/")
	if len(parts) <= 0 {
		return nil, errors.New("projectName is empty")
	}

	lastPart := parts[len(parts)-1]
	// split on capitalized characters
	words := splitOnCapitals(lastPart)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	result := strings.Join(words, "-")

	return &result, nil
}

func splitOnCapitals(s string) []string {
	var parts []string
	last := 0

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			parts = append(parts, s[last:i])
			last = i
		}
	}
	parts = append(parts, s[last:])
	return parts
}

func createMainFile(exportDir string, data pluginData) error {
	processTemplate(
		"templates/plugin-name.php.tmpl",
		filepath.Join(exportDir, data.PluginName+".php"),
		data)

	return nil
}

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
