package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
4 Битва за кондиционер
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Ошибка ввода количества тестов")
		return
	}
	var mapTestString = make([][]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		countLine, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Ошибка ввода строк в тесте данных")
			return
		}
		testString := make([]string, 0, countLine)
		for k := 0; k < countLine; k++ {
			scanner.Scan()
			text := scanner.Text()
			testString = append(testString, text)
		}
		mapTestString[i] = testString
	}

	for _, dataTest := range mapTestString {
		//fmt.Printf("Тест номер %d, количество строк %d\n", idxTest+1, len(dataTest))
		tempRange := []int{
			15,
			30,
		}
		var answer int
		for _, s := range dataTest {
			temp, err := strconv.Atoi(s[3:])
			if err == nil {
				if answer != -1 {
					if s[0:2] == ">=" {
						if temp > tempRange[1] {
							answer = -1
						}
						if temp >= tempRange[0] && temp <= tempRange[1] {
							tempRange[0] = temp
						}
					} else {
						if temp < tempRange[0] {
							answer = -1
						}
						if temp <= tempRange[1] && temp >= tempRange[0] {
							tempRange[1] = temp
						}

					}
				}
			}
			if answer != -1 {
				answer = tempRange[0]
			}
			//	fmt.Printf("Строка с показаниями t = %s\n", s)
			//fmt.Println(tempRange)
			fmt.Println(answer)

		}
		fmt.Println()
	}
}
