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
7 Печать документа
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Ошибка ввода количества тестов")
		return
	}
	var testCountPage = make([]int, t)
	var testPrintPage = make([][]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		countPage, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Ошибка ввода количества страниц в тесте")
			return
		}
		testCountPage[i] = countPage
		scanner.Scan()
		testPrintPage[i] = strings.Split(scanner.Text(), ",")
	}
	//fmt.Println("--------------------------------------")
	for idx, countPage := range testCountPage {
		printPage := testPrintPage[idx]

		arPrintPage := make([]int, 0)
		arRangePrintPage := make([][]int, 0)
		resPagePrint := make([][]int, 1, 2)
		resPagePrint[0] = []int{1, countPage}
		for _, page := range printPage {
			number, err := strconv.Atoi(page)
			if err == nil {
				//номер страницы
				arPrintPage = append(arPrintPage, number)
			} else {
				//диапазоны
				pageRange := strings.Split(page, "-")
				arRange := make([]int, 2)
				arRange[0], _ = strconv.Atoi(pageRange[0])
				arRange[1], _ = strconv.Atoi(pageRange[1])
				if arRange[0] == arRange[1] {
					arPrintPage = append(arPrintPage, arRange[0])
				} else {
					arRangePrintPage = append(arRangePrintPage, arRange)
				}

			}
		}
		//исключаем страницы
		for _, page := range arPrintPage {
			for idx, resPage := range resPagePrint {
				if len(resPage) == 0 {
					continue
				} else if len(resPage) == 1 {
					if resPage[0] == page {
						resPagePrint[idx] = []int{}
					}
				} else if page == resPage[0] || page == resPage[1] {
					if page == resPage[0] {
						resPagePrint[idx][0] = page + 1

					} else if page == resPage[1] {
						resPagePrint[idx][1] = page - 1
					}
				} else if page > resPage[0] && page < resPage[1] {
					if resPage[0] == page-1 {
						resPagePrint[idx] = []int{
							resPage[0],
						}
					} else {
						resPagePrint[idx] = []int{
							resPage[0],
							page - 1,
						}
					}
					if page+1 == resPage[1] {
						resPagePrint = append(resPagePrint, []int{
							resPage[1],
						})
					} else {
						resPagePrint = append(resPagePrint, []int{
							page + 1,
							resPage[1],
						})
					}
				}
			}

		}
		//исключаем диапазоны
		for _, rangePage := range arRangePrintPage {
			for idx, resPage := range resPagePrint {
				if len(resPage) == 0 {
					continue
				} else if len(resPage) == 1 {
					if resPage[0] >= rangePage[0] && resPage[0] <= rangePage[1] {
						resPagePrint[idx] = []int{}
					}
				} else {
					if resPage[0] > resPage[1] {
						continue
					}
					if rangePage[0] <= resPage[0] && rangePage[1] >= resPage[1] {
						// если диапазон resPage внутри rangePage
						resPagePrint[idx] = []int{}
					} else if rangePage[0] >= resPage[0] && rangePage[1] <= resPage[1] {
						// если диапазон rangePage внутри resPage
						resPagePrint[idx] = []int{
							resPage[0],
							rangePage[0] - 1,
						}
						resPagePrint = append(resPagePrint, []int{
							rangePage[1] + 1,
							resPage[1],
						})
					} else if rangePage[0] > resPage[0] && rangePage[1] < resPage[1] {
						resPagePrint[idx] = []int{
							resPage[0],
							rangePage[0] + 1,
						}
						resPagePrint = append(resPagePrint, []int{
							rangePage[1] + 1,
							resPage[1],
						})
					} else if rangePage[1] < resPage[1] && rangePage[1] > resPage[0] {
						resPagePrint[idx] = []int{
							rangePage[1] + 1,
							resPage[1],
						}
						resPagePrint = append(resPagePrint, []int{
							rangePage[1] + 1,
							resPage[0] + 1,
						})
					} else if rangePage[0] <= resPage[0] && rangePage[1] >= resPage[1] {
						resPagePrint[idx] = []int{}
					} else if rangePage[0] >= resPage[0] && rangePage[1] > resPage[1] && resPage[1] > rangePage[0] {
						resPagePrint[idx] = []int{
							resPage[0],
							rangePage[0] - 1,
						}
					} else if rangePage[0] < resPage[0] && rangePage[1] >= resPage[0] && rangePage[1] < resPage[1] {
						resPagePrint[idx] = []int{
							rangePage[1] + 1,
							resPage[1],
						}
					} else if rangePage[0] == resPage[1] {
						resPagePrint[idx][1] = resPage[1] - 1
					}
				}
			}
		}

		var s string
		for _, page := range resPagePrint {
			if len(page) == 0 {
				continue
			} else if len(page) == 1 {
				s += "," + strconv.Itoa(page[0])
			} else {
				if page[0] > page[1] {
					continue
				}
				if page[0] != page[1] {
					s += "," + strconv.Itoa(page[0]) + "-" + strconv.Itoa(page[1])
				} else {
					s += "," + strconv.Itoa(page[0])
				}
			}
		}
		fmt.Println(s[1:])
	}
}
