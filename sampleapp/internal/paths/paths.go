package paths

import (
	"fmt"

	// internal dependencies
	"/internal/config"
)

// OutputFilePath returns the path to the file where the output report is generated
func OutputFilePath(appConfig config.SampleAppConfig) string {
	var outputFilePath = fmt.Sprintf("%s/%s",
		appConfig.OutputInfo.Workspace,
		appConfig.OutputInfo.OutputFile)

	return outputFilePath
}

// TemplatePath returns the path to the template file used to generate the report
func TemplatePath(appConfig config.SampleAppConfig) string {
	templateName := "appreport.html.tmpl"

	var templatePath = fmt.Sprintf("%s/templates/%s",
		appConfig.ConfigDir, templateName)

	return templatePath
}
