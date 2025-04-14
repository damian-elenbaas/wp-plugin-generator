package generator

import (
	"errors"
	"path/filepath"
	"strings"
)

type pluginData struct {
	PluginName      string
	Slug            string
	Description     string
	Version         string
	Author          string
	TextDomain      string
	NamespaceName   string
	ClassPrefix     string
	ConstantPrefix  string
	FunctionPostfix string
}

func generatePluginData(projectName string) (*pluginData, error) {

	pluginName, err := generatePluginName(projectName)
	if err != nil {
		return nil, err
	}

	slug, err := generateSlug(projectName)
	if err != nil {
		return nil, err
	}

	classPrefix, err := generateClassPrefix(projectName)
	if err != nil {
		return nil, err
	}

	namespace, err := generateNamespace(projectName)
	if err != nil {
		return nil, err
	}

	constantPrefix, err := generateConstantPrefix(projectName)
	if err != nil {
		return nil, err
	}

	functionPostfix, err := generateFunctionPostfix(projectName)
	if err != nil {
		return nil, err
	}

	return &pluginData{
		PluginName:      *pluginName,
		Slug:            *slug,
		Description:     "A WordPress plugin",
		Version:         "0.1.0",
		Author:          "Your Name",
		TextDomain:      *slug,
		NamespaceName:   *namespace,
		ClassPrefix:     *classPrefix,
		ConstantPrefix:  *constantPrefix,
		FunctionPostfix: *functionPostfix,
	}, nil
}

// Generates the class prefix for the given projectname
func generateClassPrefix(projectName string) (*string, error) {
	lastPart, err := getLastPartOfProjectName(projectName)
	if err != nil {
		return nil, err
	}

	words := strings.Split(*lastPart, "-")
	for i := range words {
		word := words[i]
		words[i] = strings.ToUpper(word[0:1]) + word[1:]
	}
	result := strings.Join(words, "")

	return &result, nil
}

func generateNamespace(projectName string) (*string, error) {
	parts := strings.Split(projectName, "/")
	if len(parts) <= 0 {
		return nil, errors.New("projectName is empty")
	}

	for partIndex := range parts {
		words := strings.Split(parts[partIndex], "-")
		for wordIndex := range words {
			word := words[wordIndex]
			words[wordIndex] = strings.ToUpper(word[0:1]) + word[1:]
		}

		parts[partIndex] = strings.Join(words, "")
	}

	result := strings.Join(parts, "\\")

	return &result, nil
}

func generateConstantPrefix(projectName string) (*string, error) {
	lastPart, err := getLastPartOfProjectName(projectName)
	if err != nil {
		return nil, err
	}

	words := splitOnCapitals(*lastPart)
	toUpperWords(&words)
	*lastPart = strings.Join(words, "_")

	words = strings.Split(*lastPart, "-")
	toUpperWords(&words)
	result := strings.Join(words, "_")

	return &result, nil
}

func generateFunctionPostfix(projectName string) (*string, error) {
	lastPart, err := getLastPartOfProjectName(projectName)
	if err != nil {
		return nil, err
	}

	words := splitOnCapitals(*lastPart)
	toLowerWords(&words)
	*lastPart = strings.Join(words, "_")

	words = strings.Split(*lastPart, "-")
	toLowerWords(&words)
	result := strings.Join(words, "_")

	return &result, nil
}

func generatePluginName(projectName string) (*string, error) {
	lastPart, err := getLastPartOfProjectName(projectName)
	if err != nil {
		return nil, err
	}

	words := splitOnCapitals(*lastPart)
	capitalizeFirstLetterWords(&words)
	*lastPart = strings.Join(words, " ")

	words = strings.Split(*lastPart, "-")
	capitalizeFirstLetterWords(&words)
	result := strings.Join(words, " ")

	return &result, nil
}

func generateSlug(projectName string) (*string, error) {
	lastPart, err := getLastPartOfProjectName(projectName)
	if err != nil {
		return nil, err
	}

	words := splitOnCapitals(*lastPart)
	toLowerWords(&words)
	result := strings.Join(words, "-")

	return &result, nil
}

func createPluginRootFile(exportDir string, data pluginData) error {
	processTemplate(
		"plugin-name.php.tmpl",
		filepath.Join(exportDir, data.Slug+".php"),
		data)

	return nil
}

func createMainFile(exportDir string, data pluginData) error {
	processTemplate(
		"includes/PluginName.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+".php"),
		data)
	return nil
}
