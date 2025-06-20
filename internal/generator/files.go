package generator

import (
	"path/filepath"
)

func (data *pluginData) createFiles(exportDir string) error {
	data.createActivatorFile(exportDir)
	data.createAdminFiles(exportDir)
	data.createDeactivatorFile(exportDir)
	data.createGitIgnore(exportDir)
	data.createI18nFile(exportDir)
	data.createIndexFiles(exportDir)
	data.createLoaderFile(exportDir)
	data.createMainFile(exportDir)
	data.createPluginRootFile(exportDir)
	data.createPublicFiles(exportDir)
	data.createApiFiles(exportDir)

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
		"includes/PluginName_I18n.php.tmpl",
		filepath.Join(exportDir, "includes/"+data.ClassPrefix+"_I18n.php"),
		*data)
}

func (data *pluginData) createGitIgnore(exportDir string) {
	processTemplate(
		".gitignore.tmpl",
		filepath.Join(exportDir, ".gitignore"),
		*data)
}

func (data *pluginData) createAdminFiles(exportDir string) {
	processTemplate(
		"admin/PluginName_Admin.php.tmpl",
		filepath.Join(exportDir, "admin/"+data.ClassPrefix+"_Admin.php"),
		*data)
	processTemplate(
		"admin/css/plugin-name-admin.css.tmpl",
		filepath.Join(exportDir, "admin/css/"+data.Slug+"-admin.css"),
		*data)
	processTemplate(
		"admin/js/plugin-name-admin.js.tmpl",
		filepath.Join(exportDir, "admin/js/"+data.Slug+"-admin.js"),
		*data)
}

func (data *pluginData) createPublicFiles(exportDir string) {
	processTemplate(
		"public/PluginName_Public.php.tmpl",
		filepath.Join(exportDir, "public/"+data.ClassPrefix+"_Public.php"),
		*data)
	processTemplate(
		"public/css/plugin-name-public.css.tmpl",
		filepath.Join(exportDir, "public/css/"+data.Slug+"-public.css"),
		*data)
	processTemplate(
		"public/js/plugin-name-public.js.tmpl",
		filepath.Join(exportDir, "public/js/"+data.Slug+"-public.js"),
		*data)
}

func (data *pluginData) createApiFiles(exportDir string) {
	processTemplate(
		"api/PluginName_API.php.tmpl",
		filepath.Join(exportDir, "api/"+data.ClassPrefix+"_API.php"),
		*data)
}

func (data *pluginData) createIndexFiles(exportDir string) {
	for _, dir := range directories {
		processTemplate(
			"index.php.tmpl",
			filepath.Join(exportDir, dir, "index.php"),
			*data)
	}
}
