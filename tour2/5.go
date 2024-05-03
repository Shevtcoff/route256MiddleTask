package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Path struct {
	Dir     string
	Files   []string
	Folders []Path
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscanln(in, &t)
	var mapTest = make([]bytes.Buffer, t)
	for i := 0; i < t; i++ {
		var countStr int
		fmt.Fscanln(in, &countStr)
		var b bytes.Buffer
		for k := 0; k < countStr; k++ {
			text, _ := in.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			b.Write([]byte(text))
		}
		mapTest[i] = b
	}

	for _, jsonBlob := range mapTest {
		var path Path

		err := json.Unmarshal(jsonBlob.Bytes(), &path)
		if err != nil {
			fmt.Fprintln(out, "error:", err)
		}
		var count int
		searchFileInPath(path, false, &count)
		fmt.Fprintln(out, count)

	}
}
func searchFileInPath(path Path, infected bool, count *int) {
	if len(path.Files) > 0 {
		if infected {
			*count += len(path.Files)
		} else {
			for _, fileName := range path.Files {
				if strings.HasSuffix(fileName, ".hack") {
					infected = true
					*count += len(path.Files)
					break
				}
			}
		}
	}
	if len(path.Folders) > 0 {
		for _, pathInc := range path.Folders {
			searchFileInPath(pathInc, infected, count)
		}
	}
}
