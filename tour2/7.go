package main

import (
	"bufio"
	"bytes"
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
	var newLogins = make([][]byte, t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		newLogins[i] = []byte(s)
	}
	fmt.Fscan(in, &t)
	var oldLogins = make([][]byte, t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		oldLogins[i] = []byte(s)
	}

	res := make([]int, len(oldLogins))
	for idx, loginOld := range oldLogins {
		for _, loginNew := range newLogins {
			if len(loginOld) == len(loginNew) {
				if bytes.Equal(loginNew, loginOld) {
					res[idx] = 1
					break
				} else {
					subLen := len(loginOld)/2 - 1
					if subLen <= 2 || (bytes.Equal(loginOld[0:subLen], loginNew[0:subLen]) || bytes.Equal(loginOld[len(loginOld)-subLen:], loginNew[len(loginNew)-subLen:])) {
						var b bytes.Buffer
						prevSymbol := loginOld[0]
						for i := 1; i < len(loginOld); i++ {
							b.Reset()
							b.Write(loginOld[0 : i-1])
							b.Write([]byte{loginOld[i], prevSymbol})
							b.Write(loginOld[i+1:])
							if bytes.Equal(b.Bytes(), loginNew) {
								res[idx] = 1
								break
							}
							prevSymbol = loginOld[i]
						}
					}
				}
			}
		}
	}
	for _, val := range res {
		fmt.Fprintln(out, val)
	}
}
