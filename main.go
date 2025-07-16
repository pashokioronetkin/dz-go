package main

import (
	"fmt"
)

const USDtoEURO float64 = 0.86
const USDtoRUB float64 = 77.89
const RUBtoUSD float64 = 0.013
const RUBtoEURO float64 = 0.011
const EUROtoUSD float64 = 1.16
const EUROtoRUB float64 = 91.15

func getInputCurrency() string {
	var inputCurrency string
	fmt.Print("Введите исходную валюту (USD, EURO, RUB): ")
	for {
		fmt.Scan(&inputCurrency)

		if inputCurrency == "RUB" || inputCurrency == "USD" || inputCurrency == "EURO" {
			return inputCurrency
		}
		fmt.Print("Неверно выбрана валюта\nВведите заново: ")
	}
}

func getOutputCurrency(inputCurrency string) string {
	var outputCurrency string
	for {
		if inputCurrency == "RUB" {
			fmt.Print("Введите целевую валюту (USD, EURO): ")
			fmt.Scan(&outputCurrency)
			if outputCurrency == "USD" || outputCurrency == "EURO" {
				return outputCurrency
			}
		} else if inputCurrency == "USD" {
			fmt.Print("Введите целевую валюту (RUB, EURO): ")
			fmt.Scan(&outputCurrency)
			if outputCurrency == "RUB" || outputCurrency == "EURO" {
				return outputCurrency
			}
		} else if inputCurrency == "EURO" {
			fmt.Print("Введите целевую валюту (USD, RUB): ")
			fmt.Scan(&outputCurrency)
			if outputCurrency == "USD" || outputCurrency == "RUB" {
				return outputCurrency
			}
		}
	}
}

func getQuantity() float64 {
	var inputCurrencyQuantity float64
	for {
		fmt.Print("Введите количество исходной валюты: ")
		_, err := fmt.Scan(&inputCurrencyQuantity)
		if err != nil {
			fmt.Println("Ошибка: введено не число! Введите количество заново")
			var discard string
			fmt.Scanln(&discard)
			continue
		} else {
			return inputCurrencyQuantity
		}
	}
}

func calculate(currency1 string, currency2 string, quantity float64) float64 {
	fmt.Println("Конвертация из", currency1, "в", currency2)

	if currency1 == "USD" && currency2 == "EURO" {
		return USDtoEURO * quantity
	} else if currency1 == "USD" && currency2 == "RUB" {
		return USDtoRUB * quantity
	}

	if currency1 == "RUB" && currency2 == "USD" {
		return RUBtoUSD * quantity
	} else if currency1 == "RUB" && currency2 == "EURO" {
		return RUBtoEURO * quantity
	}

	if currency1 == "EURO" && currency2 == "RUB" {
		return EUROtoRUB * quantity
	} else if currency1 == "EURO" && currency2 == "USD" {
		return EUROtoUSD * quantity
	} else {
		return 0
	}
}

func main() {
	inputCurrency := getInputCurrency()
	outputCurrency := getOutputCurrency(inputCurrency)

	result := calculate(inputCurrency, outputCurrency, getQuantity())
	fmt.Print(result)
}
