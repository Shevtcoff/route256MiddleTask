package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
2 Проверка даты
*/

func main() {
	var t int
	fmt.Scan(&t)
	scanner := bufio.NewScanner(os.Stdin)
	var mapTestString = make([]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		mapTestString[i] = scanner.Text()
	}

	var mapCountDayInMonth = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	for _, s := range mapTestString {
		strDate := strings.Split(strings.Trim(s, "\n"), " ")

		day, okDay := strconv.Atoi(strDate[0])
		month, okMonth := strconv.Atoi(strDate[1])
		year, okYear := strconv.Atoi(strDate[2])
		var leapYear = false
		var currentDate = true

		if okDay != nil || okMonth != nil || okYear != nil {
			currentDate = false
		}

		if year < 1950 || year > 2300 {
			currentDate = false
		}

		if currentDate && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
			leapYear = true
		}

		if currentDate && (month < 1 || month > 12) {
			currentDate = false
		}

		if currentDate && (day < 1 || day > 31) {
			currentDate = false
		}

		if currentDate {
			maxDay, _ := mapCountDayInMonth[month]
			if month == 2 && leapYear {
				maxDay++
			}
			if day > maxDay {
				currentDate = false
			}
		}

		if currentDate {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
