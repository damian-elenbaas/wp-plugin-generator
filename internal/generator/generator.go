package generator

func GeneratePlugin(exportDir string, projectName string) error {

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
