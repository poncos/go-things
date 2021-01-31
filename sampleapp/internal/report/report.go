package report

import (
	"os"
	"text/template"
	"time"

	// internal dependencies
	"github.com/poncos/go-things/sampleapp/internal/utils"
	"github.com/poncos/go-things/sampleapp/internal/model"
	"github.com/poncos/go-things/sampleapp/internal/config"
	"github.com/poncos/go-things/sampleapp/internal/paths"
)

// RenderHTMLTemplate parses a template and saves result to file
func RenderHTMLTemplate(
	config config.SampleAppConfig,
	people []model.Person) {

	funcMap := template.FuncMap{
		"calculateAge": calculateAge,
		"timeToStr": timeToStr,
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

func timeToStr(dob time.Time) string {
	return dob.Format("01-02-2006")
}

func calculateAge(dob time.Time) int {
	now := time.Now()

	years := now.Year() - dob.Year()

	if now.Month() < dob.Month() {
		return years-1
	} else if now.Month() > dob.Month() {
		return years
	} else if now.Day() < dob.Day(){
		return years - 1
	} else {
		return years
	}
}
