package main

import "fmt"

const USDtoEURO float64 = 0.86
const USDtoRUB float64 = 77.89
const EUROtoRUB float64 = USDtoRUB / USDtoEURO

func getUserInput() (string, float64, string) {
	var quantity float64
	var currency1, currency2 string

	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&currency1)

	fmt.Print("Введите количество: ")
	fmt.Scan(&quantity)

	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&currency2)

	return currency1, quantity, currency2
}

func calculate(currency1 string, currency2 string, quantity float64) float64 {
	return USDtoEURO / quantity
}

func main() {
	currency1, quantity, currency2 := getUserInput()

	exchange := calculate(currency1, currency2, quantity)

	fmt.Print(currency1, " ", quantity, " ", currency2, '\n')
	fmt.Printf("%.2f", exchange)
}
