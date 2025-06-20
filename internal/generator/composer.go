package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func generateComposerJson(projectName string, authors []Author) (*string, error) {
	type ComposerJson struct {
		Name       string            `json:"name"`
		Type       string            `json:"type"`
		Autoload   map[string]any    `json:"autoload"`
		Authors    []Author          `json:"authors"`
		RequireDev map[string]string `json:"require-dev"`
	}

	projectNameSlug, err := generateProjectNameSlug(projectName)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(projectName, "/")
	var processedParts []string
	for _, part := range parts {
		words := strings.Split(part, "-")
		for i := range words {
			if len(words[i]) > 0 {
				words[i] = strings.ToUpper(words[i][0:1]) + words[i][1:]
			}
		}
		processedParts = append(processedParts, strings.Join(words, ""))
	}
	namespace := strings.Join(processedParts, "\\") + "\\"

	composer := ComposerJson{
		Name: *projectNameSlug,
		Type: "project",
		Autoload: map[string]any{
			"psr-4": map[string]string{
				namespace:              "includes/",
				namespace + "Public\\": "public/",
				namespace + "Admin\\":  "admin/",
				namespace + "API\\":    "api/",
			},
		},
		Authors: authors,
		RequireDev: map[string]string{
			"php-stubs/wordpress-stubs": "^6.8",
		},
	}

	jsonData, err := json.MarshalIndent(composer, "", "    ")
	if err != nil {
		return nil, err
	}

	jsonDataString := string(jsonData)

	return &jsonDataString, nil
}

func writeComposerJson(exportDir string, projectName string, authors []Author) error {
	composerJson, err := generateComposerJson(projectName, authors)
	if err != nil {
		return err
	}

	composerBytes := []byte(*composerJson)

	composerPath := filepath.Join(exportDir, "/composer.json")
	err = os.WriteFile(composerPath, composerBytes, 0644)
	if err != nil {
		log.Fatalf("Error writing composer.json: %v", err)
	}

	return nil
}

func generateProjectNameSlug(projectName string) (*string, error) {
	parts := strings.Split(projectName, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("Expecting projectname format {company}/{pluginname}, got %s", projectName)
	}

	for i := range parts {
		part := parts[i]
		words := splitOnCapitals(part)
		toLowerWords(&words)
		parts[i] = strings.Join(words, "-")
	}

	result := strings.Join(parts, "/")

	return &result, nil
}
