package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var mapTest = make([]string, t)
	for i := 0; i < t; i++ {
		var strCom string
		fmt.Fscan(in, &strCom)
		mapTest[i] = strCom
	}

	for _, dataTest := range mapTest {
		res := true
		prev := dataTest[0]
		for i := 1; i < len(dataTest); i++ {
			if i == 1 && prev != 'M' { //старт
				res = false
				break
			} else if i == len(dataTest)-1 && (dataTest[i] != 'D' || prev == 'D') { //остановка
				res = false
				break
			} else if prev == 'M' && dataTest[i] == 'M' { //запуск
				res = false
				break
			} else if prev == 'R' && dataTest[i] != 'C' { //отмена
				res = false
				break
			} else if prev == 'C' && dataTest[i] != 'M' { //перезапуск
				res = false
				break
			} else if prev == 'D' && dataTest[i] != 'M' { //запуск завершенной
				res = false
				break
			}
			prev = dataTest[i]
		}
		if len(dataTest) < 2 {
			res = false
		}

		if res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
