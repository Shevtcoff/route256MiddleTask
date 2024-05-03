package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

/*
3 Автомобильные номера
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

	for _, s := range mapTestString {
		pos := 0
		var newString string
		var curString bool = true
		for pos < len(s) {
			ok1, ok2 := false, false
			if pos+5 <= len(s) {
				numCar1 := s[pos : pos+5]
				if matched, _ := regexp.MatchString(`[A-Z]\d\d[A-Z]{2}`, numCar1); matched {
					pos += 5
					if len(newString) > 0 {
						newString += " " + numCar1
					} else {
						newString = numCar1
					}
					ok1 = true
				}
			}
			if pos+4 <= len(s) {
				numCar2 := s[pos : pos+4]
				if matched, _ := regexp.MatchString(`[A-Z]\d[A-Z]{2}`, numCar2); matched {
					pos += 4
					if len(newString) > 0 {
						newString += " " + numCar2
					} else {
						newString = numCar2
					}
					ok2 = true
				}
			}

			if !(ok1 || ok2) {
				curString = false
				break
			}

		}

		if curString {
			fmt.Println(newString)
		} else {
			fmt.Println("-")
		}

	}
}
