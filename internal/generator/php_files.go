package generator

import (
	"errors"
	"path/filepath"
	"strings"
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

	textDomain, err := generateTextDomain(projectName)
	if err != nil {
		return nil, err
	}

	classPrefix, err := generateClassPrefix(projectName)
	if err != nil {
		return nil, err
	}

	return &pluginData{
		PluginName:      projectName,
		Description:     "A WordPress plugin",
		Version:         "0.1.0",
		Author:          "Your Name",
		TextDomain:      *textDomain,
		NamespaceName:   "Namespace\\PackageName",
		PackageName:     "PackageName",
		ClassPrefix:     *classPrefix,
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

// Generates the class prefix for the given projectname
func generateClassPrefix(projectName string) (*string, error) {
	parts := strings.Split(projectName, "/")
	if len(parts) <= 0 {
		return nil, errors.New("projectName is empty")
	}

	lastPart := parts[len(parts)-1]
	words := strings.Split(lastPart, "-")
	for i := range words {
		word := words[i]
		words[i] = strings.ToUpper(word[0:1]) + word[1:]
	}
	result := strings.Join(words, "")

	return &result, nil
}

func createMainFile(exportDir string, data pluginData) error {
	processTemplate(
		"templates/plugin-name.php.tmpl",
		filepath.Join(exportDir, data.PluginName+".php"),
		data)

	return nil
}
