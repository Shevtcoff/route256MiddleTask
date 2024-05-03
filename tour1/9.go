package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
 9 Анализ игрового поля
*/

var logger bool = false

type GameField struct {
	X     int
	Y     int
	Lines [][]byte
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	var mapTestGames = make([]GameField, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		var x, y int
		coordFields := strings.Split(scanner.Text(), " ")
		y, _ = strconv.Atoi(coordFields[0])
		x, _ = strconv.Atoi(coordFields[1])
		game := GameField{
			X:     x,
			Y:     y,
			Lines: make([][]byte, y),
		}
		for k := 0; k < y; k++ {
			scanner.Scan()
			s := scanner.Text()
			game.Lines[k] = []byte(s)
		}

		mapTestGames[i] = game
	}

	//fmt.Println("-----------------------------")

	for idxTest, dataGame := range mapTestGames {
		if logger {
			fmt.Printf("Тест номер %d\n", idxTest+1)
		}
		findRect(dataGame)
	}
}

/*
*
ищет координаты прямоугольника по первой точке
*/

func findRect(game GameField) {
	var start time.Time
	resMap := make(map[string]int, 1000)
	start = time.Now()
	findIncRect(0, 0, game.X, game.Y, game, 0, resMap)
	resFinal := make([]int, 0, len(resMap))
	for _, val := range resMap {
		resFinal = append(resFinal, val)
	}
	sort.Ints(resFinal)
	fmt.Println(strings.Trim(fmt.Sprint(resFinal), "[]"))
	if logger {
		duration := time.Since(start)
		fmt.Println(duration.Seconds())
	}
}

func findIncRect(x, y, sizeX, sizeY int, game GameField, level int, resMap map[string]int) (int, bool) {
	var ok bool
	coord := make([]int, 0, 4)
	for y1 := y; y1 < sizeY-1; y1++ {
		prev := '.'
		for x1 := x; x1 < sizeX-1; x1++ {
			if prev == '.' && game.Lines[y1][x1] == '*' && (y1 == 0 || (y1 > 0 && game.Lines[y1-1][x1] == '.')) {
				coord = sizeRect(x1, y1, game)
				level++
				level, ok = findIncRect(coord[0]+1, coord[1]+1, coord[2], coord[3], game, level, resMap)
				if !ok {
					level--
				}
				key := fmt.Sprint([]int{coord[0], coord[1]})
				resMap[key] = level
				game.Lines[coord[1]][coord[0]] = '+'
				/*	for _, l := range game.Lines {
					fmt.Println(string(l))
				}*/
			}
			prev = rune(game.Lines[y1][x1])
		}
	}
	return level, ok
}

func sizeRect(x1, y1 int, game GameField) []int {
	var x2, y2 int
	for x2 = x1; x2 < len(game.Lines[y1]) && game.Lines[y1][x2] == '*'; x2++ {
	}
	for y2 = y1; y2 < len(game.Lines) && game.Lines[y2][x1] == '*'; y2++ {
	}
	return []int{x1, y1, x2 - 1, y2 - 1}
}
