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
8 3-Покер
*/

var arSuit = []string{
	"S", //Spades
	"C", //Clubs
	"D", //Diamonds
	"H", //	Hearts
}
var logger bool = false

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Ошибка ввода количества тестов")
		return
	}
	var mapTestPlayersSet = make([][][]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		countPlayers, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Ошибка ввода количества игроков в тесте")
			return
		}
		var players = make([][]string, countPlayers)
		mapTestPlayersSet[i] = players
		for k := 0; k < countPlayers; k++ {
			scanner.Scan()
			mapTestPlayersSet[i][k] = strings.Split(scanner.Text(), " ")
		}
	}

	//fmt.Println("-----------------------------")

	for idxTest, dataTest := range mapTestPlayersSet {
		if logger {
			fmt.Printf("Тест номер %d\n", idxTest+1)
		}
		var countCardUse int
		//используемые карты мапа[номер карты][масть]
		useCard := make(map[string]map[string]struct{}, len(dataTest))
		for _, cardsPlayer := range dataTest {
			for _, card := range cardsPlayer {
				number := string(card[0])
				suit := string(card[1])
				if useCard[number] == nil {
					useCard[number] = make(map[string]struct{}, 4)
				}
				useCard[number][suit] = struct{}{}
				countCardUse++
			}
		}
		//~используемые карты

		var resCars = make([]string, 0)

		for {
			var loss bool
			var res string
			var typeCombPlayer1 string
			var resCombPlayer1 = make([]string, 0, 2)
			resCombPlayer1, _ = searchComb(dataTest[0], useCard) //сет победа
			countCardUse++
			if typeCombPlayer1 == "set" {
				res = resCombPlayer1[2]
				resCars = append(resCars, res)
				continue
			}
			typeCombPlayer1, raitPlayer1 := checkComp(resCombPlayer1)
			for idxPlayer := 1; idxPlayer < len(dataTest); idxPlayer++ {
				cardsPlayer := dataTest[idxPlayer]
				cardsPlayer = append(cardsPlayer, resCombPlayer1[2])
				typeComPlayer, raitPlayer := checkComp(cardsPlayer)
				if logger {
					fmt.Println("Игрок 1", typeCombPlayer1, raitPlayer1, "Другой игрок", typeComPlayer, raitPlayer, cardsPlayer)
				}
				if typeComPlayer == "set" {
					loss = true
				} else if typeComPlayer == "double" && typeCombPlayer1 == "seniorCard" {
					loss = true
				} else if typeCombPlayer1 == typeComPlayer && raitPlayer > raitPlayer1 {
					loss = true
				}
			}
			if !loss {
				if logger {
					fmt.Println("------------win up--------------")
				}
				res = resCombPlayer1[2]
				resCars = append(resCars, res)
			}
			if countCardUse == 52 {
				break
			}
			if loss {
				continue
			}

		}

		if len(resCars) > 0 {
			fmt.Println(len(resCars))
			for _, card := range resCars {
				if logger {
					fmt.Printf("Победа игрока 1 тип комбинации: %s\n", card)
				} else {
					fmt.Println(card)
				}
			}
		}

		if len(resCars) == 0 {
			fmt.Println(0)
		}
	}
}

func checkComp(cardsPlayer []string) (string, int) {
	var typeComb string
	var rait int
	var rating = map[string]int{
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	number1 := cardsPlayer[0][0]
	number2 := cardsPlayer[1][0]
	number3 := cardsPlayer[2][0]
	switch {
	case number1 == number2 && number2 == number3:
		typeComb = "set"
	case number1 == number2 || number2 == number3 || number1 == number3:
		typeComb = "double"
		var val string
		var ok bool
		if number1 == number2 || number1 == number3 {
			val = string(number1)
		} else {
			val = string(number2)
		}
		if rait, ok = rating[val]; !ok {
			rait, _ = strconv.Atoi(val)
		}
	default:
		typeComb = "seniorCard"
		for i := 0; i < len(cardsPlayer); i++ {
			card := cardsPlayer[i]
			var ok bool
			var val int
			if val, ok = rating[string(card[0])]; !ok {
				val, _ = strconv.Atoi(string(card[0]))
			}
			if val > rait {
				rait = val
			}
		}
	}

	return typeComb, rait
}

func searchComb(cardsPlayer []string, useCard map[string]map[string]struct{}) ([]string, string) {
	var typeComb string
	var rating = map[int]string{
		10: "T",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
	}
	var resComb = cardsPlayer
	number1 := string(cardsPlayer[0][0])
	number2 := string(cardsPlayer[1][0])
	var resNumber string
	var resSuit string
	switch {
	case number1 == number2 && len(useCard[number1]) < 4:
		typeComb = "set"
		resNumber = number1
		for _, suit := range arSuit {
			if _, ok := useCard[resNumber][suit]; !ok {
				resSuit = suit
				break
			}
		}
		resComb = append(resComb, resNumber+resSuit)
		if logger {
			fmt.Printf("найдена комбинация сет %s\n", resComb)
		}
	case len(useCard[number1]) == 4 && number1 == number2:
		typeComb = "double"
		var findCard bool
		var i int
		for i = 2; i < 15; i++ {
			ind := strconv.Itoa(i)
			if i > 9 {
				ind = rating[i]
			}
			for _, suit := range arSuit {
				if _, ok := useCard[ind][suit]; !ok {
					resSuit = suit
					if i < 10 {
						resNumber = ind
					} else {
						resNumber = rating[i]
					}
					findCard = true
					break
				}
			}
			if findCard {
				break
			}
		}
		resComb = append(resComb, resNumber+resSuit)
		if logger {
			fmt.Printf("найдена комбинация дубль 1 и 2 %s\n", resComb)
		}
	case number1 != number2 && len(useCard[number1]) < 4:
		typeComb = "double"
		resNumber = number1
		for _, suit := range arSuit {
			if _, ok := useCard[number1][suit]; !ok {
				resSuit = suit
				break
			}
		}
		resComb = append(resComb, resNumber+resSuit)
		if logger {
			fmt.Printf("найдена комбинация дубль 1 и 3 %s\n", resComb)
		}
	case number1 != number2 && len(useCard[number2]) < 4:
		typeComb = "double"
		resNumber = number2
		for _, suit := range arSuit {
			if _, ok := useCard[number2][suit]; !ok {
				resSuit = suit
				break
			}
		}
		resComb = append(resComb, resNumber+resSuit)
		if logger {
			fmt.Printf("найдена комбинация дубль 2 и 3 %s\n", resComb)
		}
	default:
		typeComb = "seniorCard"
		var i int
	out:
		for i = 2; i < 15; i++ {
			resNumber = strconv.Itoa(i)
			if i > 9 {
				resNumber = rating[i]
			}
			for _, suit := range arSuit {
				if _, ok := useCard[resNumber][suit]; !ok {
					resSuit = suit
					break out
				}
			}
		}
		resComb = append(resComb, resNumber+resSuit)
		if logger {
			fmt.Printf("найдена старшая карта %s\n", resComb)
		}
	}
	if useCard[resNumber] == nil {
		useCard[resNumber] = make(map[string]struct{}, 4)
	}
	useCard[resNumber][resSuit] = struct{}{}
	return resComb, typeComb
}
