package generator

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GeneratePlugin(exportDir string, projectName string) error {
	if containsInvalidCharacters(projectName) {
		return errors.New("project name contains invalid characters")
	}

	pluginData, err := generatePluginData(projectName)
	if err != nil {
		return err
	}

	err = createDirectories(exportDir)
	if err != nil {
		return err
	}

	err = writeComposerJson(exportDir, projectName, []Author{})
	if err != nil {
		return err
	}

	err = pluginData.createFiles(exportDir)
	if err != nil {
		return err
	}

	return nil
}

// check if given string contains any invalidcharacter
func containsInvalidCharacters(s string) bool {
	invalidChars := []string{".", "'", ","}
	for _, char := range invalidChars {
		if strings.Contains(s, char) {
			return true
		}
	}
	return false
}

var directories = []string{
	"",
	"includes",
	"public",
	"public/css",
	"public/js",
	"admin",
	"admin/css",
	"admin/js",
	"languages",
	"vendor",
}

func createDirectories(exportDir string) error {
	// Create directories
	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(exportDir, dir), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
