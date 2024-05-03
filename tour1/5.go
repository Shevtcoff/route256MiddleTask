package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
*
5 Компрессия данных
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Ошибка ввода количества тестов")
		return
	}
	var mapTestNumber = make([][]int, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		countNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Ошибка ввода строк в тесте данных")
			return
		}
		scanner.Scan()
		arNumber := strings.Split(scanner.Text(), " ")
		testNumber := make([]int, 0, countNumber)
		for _, val := range arNumber {
			number, err := strconv.Atoi(val)
			if err == nil {
				testNumber = append(testNumber, number)
			}
		}
		mapTestNumber[i] = testNumber
	}

	for _, dataTest := range mapTestNumber {
		resNumber1 := compressionVar1(dataTest)
		resNumber2 := compressionVar2(dataTest)
		if len(resNumber1) > len(resNumber2) {
			fmt.Println(len(resNumber2))
			fmt.Println(strings.Trim(fmt.Sprint(resNumber2), "[]"))
		} else {
			fmt.Println(len(resNumber1))
			fmt.Println(strings.Trim(fmt.Sprint(resNumber1), "[]"))
		}
		//fmt.Printf("Тест номер %d, количество чисел %d\n", idxTest+1, len(dataTest))
		//fmt.Printf("Последовательность %v компрессия %v\n", dataTest, resNumber1)
		//fmt.Printf("Последовательность %v компрессия %v\n", dataTest, resNumber2)

	}
}

func compressionVar1(arNumber []int) []int {
	resNumber := make([]int, 0)
	prev := arNumber[0]
	resNumber = append(resNumber, prev)
	var typeSeq, count int
	for i := 1; i < len(arNumber); i++ {
		if typeSeq <= 0 && prev == arNumber[i]-1 {
			typeSeq = -1
			count++
		} else if typeSeq >= 0 && prev == arNumber[i]+1 {
			typeSeq = 1
			count--
		} else {
			typeSeq = 0
		}
		if typeSeq == 0 {
			resNumber = append(resNumber, count, arNumber[i])
			count = 0
		}
		prev = arNumber[i]
	}
	resNumber = append(resNumber, count)
	return resNumber
}

func compressionVar2(arNumber []int) []int {
	resNumber := make([]int, 0)
	prev := arNumber[0]
	resNumber = append(resNumber, prev)
	var typeSeq, count int
	for i := 1; i < len(arNumber); i++ {
		if typeSeq <= 0 && prev == arNumber[i]-1 {
			typeSeq = -1
			count++
		} else if typeSeq >= 0 && prev == arNumber[i]+1 && prev > 0 {
			typeSeq = 1
			count--
		} else {
			typeSeq = 0
		}
		if typeSeq == 0 {
			resNumber = append(resNumber, count, arNumber[i])
			count = 0
		}
		prev = arNumber[i]
	}
	resNumber = append(resNumber, count)
	return resNumber
}
