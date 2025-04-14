package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func GeneratePlugin(exportDir string, projectName string) error {

	err := writeComposerJson(exportDir, projectName, []Author{})
	if err != nil {
		return err
	}

	pluginData := generatePluginData(projectName)

	err = createMainFile(exportDir, *pluginData)
	if err != nil {
		return err
	}

	return nil
}

type PluginData struct {
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

func generatePluginData(projectName string) *PluginData {
	return &PluginData{
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
	}
}

func createMainFile(exportDir string, data PluginData) error {
	processTemplate("templates/plugin-name.php.tmpl", filepath.Join(exportDir, data.PluginName+".php"), data)

	return nil
}

func processTemplate(templatePath, outputPath string, data PluginData) {
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
