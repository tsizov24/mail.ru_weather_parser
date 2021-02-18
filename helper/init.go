package helper

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func Init() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: parseWeather,
	}).Start()
}

func parseWeather(_ *geziyor.Geziyor, r *client.Response) {
	getWeatherNow(r.HTMLDoc)
	getWeatherPeriod(r.HTMLDoc)
	getWeatherDay(r.HTMLDoc)
}
