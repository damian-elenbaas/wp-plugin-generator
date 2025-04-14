package generator

import (
	"errors"
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

	err = writeComposerJson(exportDir, projectName, []Author{})
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
