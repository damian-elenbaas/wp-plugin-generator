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

	err = createPluginRootFile(exportDir, *pluginData)
	if err != nil {
		return err
	}

	err = createMainFile(exportDir, *pluginData)
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

func createDirectories(exportDir string) error {
	// Create root directories if they don't exist
	err := os.MkdirAll(exportDir, 0755)
	if err != nil {
		return err
	}

	// Create includes directory
	err = os.MkdirAll(filepath.Join(exportDir, "includes"), 0755)
	if err != nil {
		return err
	}

	// Create public directory
	err = os.MkdirAll(filepath.Join(exportDir, "public"), 0755)
	if err != nil {
		return err
	}

	// Create admin directory
	err = os.MkdirAll(filepath.Join(exportDir, "admin"), 0755)
	if err != nil {
		return err
	}

	// Create languages directory
	err = os.MkdirAll(filepath.Join(exportDir, "languages"), 0755)
	if err != nil {
		return err
	}

	return nil
}
