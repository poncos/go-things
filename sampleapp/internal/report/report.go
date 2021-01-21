package templates

import (
	"github.com/poncos/go-things/sampleapp/internal/utils"
	// internal dependencies
	"github.com/poncos/go-things/samplemodule/internal/model"
	"github.com/poncos/go-things/samplemodule/internal/config"
)

// RenderHTMLTemplate parses a template and saves result to file
func RenderHTMLTemplate(
	config config.SampleAppConfig,
	people []model.Person) {

	funcMap := template.FuncMap{
		"calculateAge": calculateAge,
	}

	templateName := "appreport.html.tmpl"
	path := config.ConfigDir + "/templates/" + templateName

	tmpl, err := template.New(templateName).Funcs(funcMap).ParseFiles(path)

	if err != nil {
		utils.Fatal("ERROR creating template: %v", err)
	}

	err = tmpl.Execute(f, people)
	if err != nil {
		utils.Fatal("ERROR executing template: %s", err)
		return
	}
}

func calculateAge(dob Time) string {
	return dob.Format("01-02-2006")
}
