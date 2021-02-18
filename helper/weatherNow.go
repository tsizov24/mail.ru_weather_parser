package helper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var (
	nowTemperature      int64
	nowWeatherCondition string
	nowFeelsTemperature int64
	nowHumidity         int64
	nowPressure         int64
	nowUltravioletIndex int64
	nowWindSpeed        int64
)

func getWeatherNow(doc *goquery.Document) {
	nowTemperature = strToNum(cutFrom(cutSpaces(doc.Find("div.information__content__temperature").Text()), "°"))
	doc.Find("div.information__content__additional__item").Each(func(i int, s *goquery.Selection) {
		switch i {
			case 0: nowFeelsTemperature = strToNum(cutFrom(cutSpaces(s.After("span.information__content__additional__item__info").Text()), "°"))
			case 1: nowWeatherCondition = cutSpaces(s.Text())
			case 2: nowPressure = strToNum(cutFrom(cutSpaces(s.Text()), " "))
			case 3: nowHumidity = strToNum(cutFrom(cutSpaces(s.Text()), "%"))
			case 4: nowWindSpeed = strToNum(cutFrom(cutSpaces(s.Text()), " "))
			case 5: nowUltravioletIndex = strToNum(cutSpaces(s.Text()))
		}
	})
}

func printWeatherNow() {
	fmt.Println("\tСейчас")
	fmt.Printf("%s: %d\n", "Температура", nowTemperature)
	fmt.Printf("%s: %d\n", "Ощущается как", nowFeelsTemperature)
	fmt.Printf("%s: %s\n", "Погодные условия", nowWeatherCondition)
	fmt.Printf("%s: %d\n", "Атмосферное давление", nowPressure)
	fmt.Printf("%s: %d\n", "Влажность", nowHumidity)
	fmt.Printf("%s: %d\n", "Скорость ветра", nowWindSpeed)
	fmt.Printf("%s: %d\n", "Индекс ультрафиолета", nowUltravioletIndex)
	fmt.Println()
}
