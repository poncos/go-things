package main

import (
	"log"
	"time"

	"github.com/poncos/gothings/sampleapp/internal/config"
	"github.com/poncos/gothings/sampleapp/internal/model"
	"github.com/poncos/gothings/sampleapp/internal/report"
)

func main() {
	var people = make([]model.Person, 10)
	var appConfig = config.LoadConfig()

	for index := range people {
		location, err := time.LoadLocation("Spain/Madrid")

		if err != nil {
			log.Fatal("Error creation location variable", err)
		}

		people[index] = model.Person{
			"David",
			"Column",
			index % 2,
			time.Date(1981, 8, 10, 18, 30, 0, 0, location)}
	}

	report.RenderHTMLTemplate(appConfig, people)
}
