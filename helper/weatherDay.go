package helper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var (
	dayTitle                    [10]string
	dayWeatherCondition         [10]string
	dayTemperatureDay           [10]int64
	dayTemperatureNight         [10]int64
	dayPressure                 [10]int64
	dayHumidity                 [10]int64
	dayWindSpeed                [10]int64
	dayUltravioletIndex         [10]int64
	dayPrecipitationProbability [10]int64
)

func getWeatherDay(doc *goquery.Document) {
	doc.Find("div.day_index").Each(func(i int, s *goquery.Selection) {
		s = s.First()
		dayTitle[i] = s.Find("div").First().Text()
		dayWeatherCondition[i] = s.Find("div.weather-icon").AttrOr("title", "!")
		dayTemperatureDay[i] = strToNum(cutFrom(cutSpaces(s.Find("div.day__temperature").Text()), "°"))
		dayTemperatureNight[i] = strToNum(cutFrom(cutSpaces(s.Find("span.day__temperature__night").Text()), "°"))
		s.Find("div.day__additional").Each(func(i2 int, s2 *goquery.Selection) {
			switch i2 {
				case 0: dayPressure[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), " "))
				case 1: dayHumidity[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), "%"))
				case 2: dayWindSpeed[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), " "))
				case 3: dayUltravioletIndex[i] = strToNum(cutSpaces(s2.Find("span").After("icon").Text()))
				case 4: dayPrecipitationProbability[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), "%"))
			}
		})
	})
}

func printWeatherDay() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\t%s\n", dayTitle[i])
		fmt.Printf("%s: %s\n", "Погодные условия", dayWeatherCondition[i])
		fmt.Printf("%s: %d\n", "Температура днем", dayTemperatureDay[i])
		fmt.Printf("%s: %d\n", "Температура ночью", dayTemperatureNight[i])
		fmt.Printf("%s: %d\n", "Атмосферное давление", dayPressure[i])
		fmt.Printf("%s: %d\n", "Влажность", dayHumidity[i])
		fmt.Printf("%s: %d\n", "Скорость ветра", dayWindSpeed[i])
		fmt.Printf("%s: %d\n", "Индекс ультрафиолета", dayUltravioletIndex[i])
		fmt.Printf("%s: %d\n", "Вероятность осадков", dayPrecipitationProbability[i])
		fmt.Println()
	}
}
