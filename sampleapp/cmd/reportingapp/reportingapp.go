package main

import (
	"time"
	"log"

	"github.com/poncos/go-things/sampleapp/internal/model"
	"github.com/poncos/go-things/sampleapp/internal/report"
	"github.com/poncos/go-things/sampleapp/internal/config"
)

func main() {
	var people [100]model.Person
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
			time.Date(1981, 8, 10, 18, 30, 0, 0, location)
		}
	}

	report.RenderHTMLTemplate(appConfig, people)
}
