package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserOperation() string {
	var operation string
	fmt.Print("Введите операцию, которую хотите выполнить (AVG, SUM, MED): ")
	for {
		fmt.Scan(&operation)

		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		}
		fmt.Print("Такой операции нет\nВведите заново: ")
	}
}

// func sliceFill(array []int) []int {
//   var n int
//   var choice string
//   fmt.Println("Заполнение массива")
//   for choice != "n" {
//     fmt.Print("Введите число: ")
//     fmt.Scan(&n)
//     array = append(array, n)

//     fmt.Print("Продолжить? y/n ")
//     fmt.Scan(&choice)
//     if choice != "y" && choice != "n" {
//       fmt.Print("Неправильная команда, введите еще раз: ")
//       fmt.Scan(&choice)
//     }
//   }
//   return array
// }

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите последовательность целых чисел через запятую: ")
	var cleanBuffer string
	fmt.Scanln(&cleanBuffer)
	_ = scanner.Scan()
	if len(scanner.Text()) == 0 {
		return "0"
	}
	return scanner.Text()
}

func stringToArray(s string) []int {
	var array []int
	parts := strings.Split(s, ",")

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		if n, err := strconv.Atoi(trimmed); err == nil {
			array = append(array, n)
		}
	}

	return array
}

func sumArray(array []int) float64 {
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return float64(sum)
}

func medArray(array []int) float64 {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	mid := len(array) / 2
	if len(array)%2 != 0 {
		return float64(array[mid])
	} else {
		return float64(array[mid-1]+array[mid]) / 2
	}
}

func opOnArray(operation string, array []int) float64 {
	if operation == "AVG" {
		return sumArray(array) / float64(len(array))
	} else if operation == "SUM" {
		return sumArray(array)
	} else {
		return medArray(array)
	}
}

func main() {
	operation := getUserOperation()
	userStr := getUserInput()
	slice := stringToArray(userStr)
	result := opOnArray(operation, slice)
	fmt.Println(result)
}
