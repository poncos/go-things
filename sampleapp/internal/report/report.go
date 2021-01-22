package report

import (
	"os"
	"text/template"
	"time"

	// internal dependencies
	"github.com/poncos/gothings/sampleapp/internal/utils"
	"github.com/poncos/gothings/sampleapp/internal/model"
	"github.com/poncos/gothings/sampleapp/internal/config"
	"github.com/poncos/gothings/sampleapp/internal/paths"
)

// RenderHTMLTemplate parses a template and saves result to file
func RenderHTMLTemplate(
	config config.SampleAppConfig,
	people []model.Person) {

	funcMap := template.FuncMap{
		"calculateAge": calculateAge,
	}

	path := paths.TemplatePath(config)
	outputPath := paths.OutputFilePath(config)

	tmpl, err := template.New("appreport.html.tmpl").Funcs(funcMap).ParseFiles(path)

	if err != nil {
		utils.Fatal("ERROR creating template: %v", err)
	}

	f, errCreateFile := os.Create(outputPath)

	if errCreateFile != nil {
		utils.Fatal("ERROR opening report file %s", err)
	}

	err = tmpl.Execute(f, people)
	if err != nil {
		utils.Fatal("ERROR executing template: %s", err)
		return
	}
}

func calculateAge(dob time.Time) string {
	return dob.Format("01-02-2006")
}
