package generator

func GeneratePlugin(exportDir string, projectName string) error {

	err := writeComposerJson(exportDir, projectName, []Author{})
	if err != nil {
		return err
	}

	return nil
}
