package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 1 Морской бой
*/

func main() {
	var t int
	var mapRank = map[int]int{
		1: 4,
		2: 3,
		3: 2,
		4: 1,
	}
	fmt.Scan(&t)
	scanner := bufio.NewScanner(os.Stdin)
	var arString = make([]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		arString[i] = scanner.Text()
	}
	for _, s := range arString {
		shipRank := strings.Split(strings.Trim(s, "\n"), " ")
		var arShips = make(map[int]int, 4)
		msg := "YES"
		for _, ship := range shipRank {
			numShit, err := strconv.Atoi(ship)
			if err == nil {
				arShips[numShit]++
				if arShips[numShit] > mapRank[numShit] {
					msg = "NO"
					break
				}
			}

		}
		fmt.Println(msg)
	}
}
