package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
6 Терминал
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
		res := make([]string, 1)
		var pointer int
		var line int = 0
		for _, char := range s {
			switch char {
			case 'L':
				if pointer-1 >= 0 {
					pointer--
				}
			case 'R':
				if pointer+1 <= len(res[line]) {
					pointer++
				}
			case 'U':
				if line > 0 {
					line--
					if pointer > len(res[line]) {
						pointer = len(res[line])
					}
				}
			case 'D':
				if line+1 < len(res) {
					line++
				}
				if pointer > len(res[line]) {
					pointer = len(res[line])
				}
			case 'B':
				pointer = 0
			case 'E':
				pointer = len(res[line])
			case 'N':
				res = append(res, "")
				if line < len(res) {
					for i := len(res) - 1; i > line; i-- {
						res[i] = res[i-1]
					}
					res[line+1] = ""
				}
				if pointer < len(res[line]) {
					res[line], res[line+1] = res[line][0:pointer], res[line][pointer:]
				}
				pointer = 0
				line++
			default:
				//fmt.Printf("Строка \"%s\" курсор %d символ %c\n", res[line], pointer, char)
				if pointer >= len(res[line]) {
					res[line] += string(char)
				} else if pointer == 0 {
					res[line] = string(char) + res[line]
				} else {
					res[line] = res[line][0:pointer] + string(char) + res[line][pointer:]
				}
				pointer++
			}
			//fmt.Println(res, pointer, string(char), line)
		}

		for _, s := range res {
			fmt.Println(s)
		}
		fmt.Println("-")

	}
}
