package main

import (
	"bufio"
	"fmt"
	"os"
)

type Diary struct {
	N    int
	M    int
	Eval [][]byte
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	var mapTest = make([]Diary, t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		diary := Diary{
			N:    n,
			M:    m,
			Eval: make([][]byte, n),
		}
		for k := 0; k < n; k++ {
			var s string
			fmt.Fscan(in, &s)
			diary.Eval[k] = []byte(s)
		}
		mapTest[i] = diary
	}
	fmt.Println("----------------------")

	for _, dataTest := range mapTest {

		minEval := byte('6')
		minLine, minCol := 0, 0
	loop:
		for i := 0; i < dataTest.M; i++ { //столбцы
			for k := 0; k < dataTest.N; k++ { //строки
				if dataTest.Eval[k][i] < minEval {
					minLine = k + 1
					minEval = dataTest.Eval[k][i]
					if minEval == byte('1') {
						break loop
					}
				}
			}

		}
		//	fmt.Println("Минимальная в строке", string(minEval), minLine)
		minEval = byte('6')
	loop2:
		for i := 0; i < dataTest.N; i++ { //строки
			for k := 0; k < dataTest.M; k++ { //столбцы
				if dataTest.Eval[i][k] < minEval && minLine != i+1 {
					//fmt.Println("нашел ", i, k, string(dataTest.Eval[i][k]))
					minEval = dataTest.Eval[i][k]
					minCol = k + 1
					if minEval == byte('1') {
						break loop2
					}

				}
			}
		}
		//	fmt.Println("Минимальная в столбце", string(minEval), minCol)
		for _, eval := range dataTest.Eval {
			fmt.Fprintln(out, string(eval))
		}
		fmt.Fprintln(out, minLine, minCol)

	}
}
