package generator

import (
	"path/filepath"
)

func (data *pluginData) createFiles(exportDir string) error {
	data.createPluginRootFile(exportDir)
	data.createMainFile(exportDir)
	data.createLoaderFile(exportDir)
	data.createActivatorFile(exportDir)
	data.createDeactivatorFile(exportDir)
	data.createI18nFile(exportDir)
	data.createGitIgnore(exportDir)

	return nil
}

func (data *pluginData) createPluginRootFile(exportDir string) {
	processTemplate(
		"plugin-name.php.tmpl",
		filepath.Join(exportDir, data.Slug+".php"),
		*data)
}

func (data *pluginData) createMainFile(exportDir string) {
	processTemplate(
		"includes/PluginName.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+".php"),
		*data)
}

func (data *pluginData) createLoaderFile(exportDir string) {
	processTemplate(
		"includes/PluginName_Loader.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+"_Loader.php"),
		*data)
}

func (data *pluginData) createActivatorFile(exportDir string) {
	processTemplate(
		"includes/PluginName_Activator.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+"_Activator.php"),
		*data)
}

func (data *pluginData) createDeactivatorFile(exportDir string) {
	processTemplate(
		"includes/PluginName_Deactivator.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+"_Deactivator.php"),
		*data)
}

func (data *pluginData) createI18nFile(exportDir string) {
	processTemplate(
		"includes/PluginName_i18n.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+"_i18n.php"),
		*data)
}

func (data *pluginData) createGitIgnore(exportDir string) {
	processTemplate(
		".gitignore.tmpl",
		filepath.Join(exportDir, ".gitignore"),
		*data)
}
