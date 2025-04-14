package generator

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

	// Trigger 'composer install' to download all packages and create the autoload setup.
	err = runComposerInstall(exportDir)
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

func runComposerInstall(exportDir string) error {
	// Run composer install
	cmd := exec.Command("composer", "install")
	cmd.Dir = exportDir
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running composer install: %v\n", err)
		return err
	}

	return nil
}
