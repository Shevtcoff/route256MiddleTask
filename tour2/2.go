package main

import (
	"bufio"
	"fmt"
	"os"
)

type PaymentFee struct {
	pr   int
	data []int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var mapTest = make([]PaymentFee, t)
	for i := 0; i < t; i++ {
		var count, procent int
		fmt.Fscan(in, &count, &procent)
		var data = make([]int, 0, count)
		for k := 0; k < count; k++ {
			var num int
			fmt.Fscan(in, &num)
			data = append(data, num)
		}
		mapTest[i] = PaymentFee{
			pr:   procent,
			data: data,
		}
	}

	for _, dataTest := range mapTest {
		var sum float64
		for _, price := range dataTest.data {
			sum += float64(price)*(float64(dataTest.pr)/100) - float64(price*dataTest.pr/100)
		}
		if sum < 0 {
			sum *= -1
		}
		fmt.Fprintf(out, "%.2f\n", sum)
	}

}
