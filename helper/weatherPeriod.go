package helper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var (
	periodTitle                    [2]string
	periodTemperature              [2]int64
	periodWeatherCondition         [2]string
	periodPressure                 [2]int64
	periodHumidity                 [2]int64
	periodWindSpeed                [2]int64
	periodUltravioletIndex         [2]int64
	periodPrecipitationProbability [2]int64
)

func getWeatherPeriod(doc *goquery.Document) {
	doc.Find("div.information__content__period").Each(func(i int, s *goquery.Selection) {
		periodTitle[i] = s.Find("div.information__content__period__title").Text()
		periodTemperature[i] = strToNum(cutFrom(s.Find("div.information__content__period__temperature").Text(), "°"))
		periodWeatherCondition[i] = s.Find("div.weather-icon").AttrOr("title", "")
		s.Find("div.information__content__period__additional__item").Each(func(i2 int, s2 *goquery.Selection) {
			switch i2 {
				case 0: periodPressure[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), " "))
				case 1: periodHumidity[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), "%"))
				case 2: periodWindSpeed[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), " "))
				case 3: periodUltravioletIndex[i] = strToNum(cutSpaces(s2.Find("span").After("icon").Text()))
				case 4: periodPrecipitationProbability[i] = strToNum(cutFrom(cutSpaces(s2.Find("span").After("icon").Text()), "%"))
			}
		})
	})
}

func printWeatherPeriod() {
	for i := 0; i < 2; i++ {
		fmt.Printf("\t%s\n", periodTitle[i])
		fmt.Printf("%s: %d\n", "Температура", periodTemperature[i])
		fmt.Printf("%s: %s\n", "Погодные условия", periodWeatherCondition[i])
		fmt.Printf("%s: %d\n", "Атмосферное давление", periodPressure[i])
		fmt.Printf("%s: %d\n", "Влажность", periodHumidity[i])
		fmt.Printf("%s: %d\n", "Скорость ветра", periodWindSpeed[i])
		fmt.Printf("%s: %d\n", "Индекс ультрафиолета", periodUltravioletIndex[i])
		fmt.Printf("%s: %d\n", "Вероятность осадков", periodPrecipitationProbability[i])
		fmt.Println()
	}
}
