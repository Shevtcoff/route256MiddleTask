package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	var mapTest = make([][]int, t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		var season = make([]int, n)
		for k := 0; k < n; k++ {
			var val int
			fmt.Fscan(in, &val)
			season[k] = val
		}
		mapTest[i] = season
	}
	//fmt.Println("----------------------")

	for _, dataTest := range mapTest {

		res := make([][]int, 0, len(dataTest))

		prev := 0
		vektor := 0
		subRes := make([]int, 0, len(dataTest))
		subRes = append(subRes, 0)
		for i := 1; i < len(dataTest); i++ {

			//	fmt.Fprintln(out, dataTest[i], prev)
			if (vektor == 0 || vektor == 1) && dataTest[prev] < dataTest[i] { // начало или рост
				subRes = append(subRes, i)
				vektor = 1
			} else if (vektor == -1 || vektor == 1) && dataTest[prev] > dataTest[i] { //снижение
				subRes = append(subRes, i)
				vektor = -1
			} else if vektor == -1 && dataTest[prev] < dataTest[i] { // конец сезона
				vektor = 0
				if len(subRes) > 2 {
					if len(subRes)%2 == 0 {
						res = append(res, subRes[0:len(subRes)-1])
					} else {
						res = append(res, subRes)
					}
					i--
				}
				subRes = make([]int, 0, len(dataTest))
				subRes = append(subRes, prev)
			} else { // застрой
				vektor = 0
				if len(subRes) > 2 {
					if len(subRes)%2 == 0 {
						res = append(res, subRes[0:len(subRes)-1])
					} else {
						res = append(res, subRes)
					}
					i--
				}
				subRes = make([]int, 0, len(dataTest))
				subRes = append(subRes, prev)
			}
			prev = i
		}
		if len(subRes) > 2 {
			if len(subRes)%2 == 0 {
				res = append(res, subRes[0:len(subRes)-1])
			} else {
				res = append(res, subRes)
			}

		}
		resCount := make([]int, 0, len(dataTest))
		prevLastIndx := 0
		nsSes := 1

		for _, sub := range res {
			if prevLastIndx > 0 && prevLastIndx == sub[0] {
				if nsSes == 1 {
					resCount = resCount[0 : len(resCount)-1]
				}
				nsSes++
			} else {
				if nsSes > 1 {
					resCount = append(resCount, nsSes)
				} else {
					nsSes = 1
					oneSes := false
					prev := 0
					for i := 1; i < len(sub); i++ {
						if (dataTest[prev]+1 == dataTest[i] && i <= (len(sub)/2)) || (i > (len(sub)/2) && dataTest[prev]-1 == dataTest[i]) {
							oneSes = true
						} else {
							oneSes = false
							break
						}
						prev = i
					}
					if oneSes {
						countSes := (len(sub) - 1) / 2
						for i := 0; i < countSes; i++ {
							resCount = append(resCount, 1)
						}
					} else {
						resCount = append(resCount, 1)
					}
				}
			}
			prevLastIndx = sub[len(sub)-1]
		}
		if nsSes > 1 {
			resCount = append(resCount, nsSes)
		}
		for len(resCount) < len(dataTest) {
			resCount = append(resCount, 0)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(resCount)))
		fmt.Fprintln(out, strings.Trim(fmt.Sprint(resCount), "[]"))

	}
}
