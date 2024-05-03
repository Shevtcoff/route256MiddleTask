package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storage struct {
	N      int
	M      int
	Coord  [][]byte
	coordA [2]int
	coordB [2]int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	var mapStorage = make([]Storage, t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		store := Storage{
			N:     n,
			M:     m,
			Coord: make([][]byte, n),
		}

		for k := 0; k < n; k++ {
			var s string
			fmt.Fscan(in, &s)
			store.Coord[k] = []byte(s)
			for idx, val := range s {
				if val == 'A' {
					store.coordA = [2]int{k, idx}
				} else if val == 'B' {
					store.coordB = [2]int{k, idx}
				}
			}
		}
		mapStorage[i] = store
	}

	for _, dataTest := range mapStorage {

		if dataTest.coordA[0] < dataTest.coordB[0] || (dataTest.coordA[0] == dataTest.coordB[0] && dataTest.coordA[1] < dataTest.coordB[1]) {
			//вврех А
			searchPath(byte('a'), dataTest.coordA, 1, dataTest)
			searchPath(byte('b'), dataTest.coordB, 0, dataTest)
		} else {
			//вверх B
			searchPath(byte('a'), dataTest.coordA, 0, dataTest)
			searchPath(byte('b'), dataTest.coordB, 1, dataTest)
		}
		for _, char := range dataTest.Coord {
			fmt.Fprintln(out, string(char))
		}
	}
}

func searchPath(symbol byte, coordSymbol [2]int, direction int, coord Storage) {
	posY, posX := coordSymbol[0], coordSymbol[1]
	posFinX, posFinY := 0, 0
	if direction == 0 {
		posFinX, posFinY = coord.M-1, coord.N-1
	}
	for posX != posFinX || posY != posFinY {
		if direction == 1 {
			if posY > 0 && coord.Coord[posY-1][posX] == '.' {
				posY--
			} else if posX > 0 && coord.Coord[posY][posX-1] == '.' {
				posX--
			}
		} else {
			if posY < posFinY && coord.Coord[posY+1][posX] == '.' {
				posY++
			} else if posX < posFinX && coord.Coord[posY][posX+1] == '.' {
				posX++
			}
		}
		coord.Coord[posY][posX] = symbol
	}
}
