package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

/*
10 Дерево комментариев
*/

var logger bool = false
var loggerFile bool = false

type Comment struct {
	id       int
	idParent int
	text     string
	use      bool
	level    int
}

type Post struct {
	count    int
	comments map[int]Comment
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	var mapTestPost = make([]Post, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		count, _ := strconv.Atoi(scanner.Text())
		post := Post{
			count:    count,
			comments: make(map[int]Comment, count),
		}
		for k := 0; k < count; k++ {
			scanner.Scan()
			s := scanner.Text()
			re, _ := regexp.Compile(`^(\d+)\s(-?\d+)\s(.+)?`)
			res := re.FindAllStringSubmatch(s, -1)
			if len(res) > 0 && len(res[0]) > 2 {
				id, _ := strconv.Atoi(res[0][1])
				idParent, _ := strconv.Atoi(res[0][2])
				post.comments[id] = Comment{
					id:       id,
					idParent: idParent,
					text:     res[0][3],
					use:      false,
					level:    0,
				}
			}
		}

		mapTestPost[i] = post
	}

	/*	fmt.Println("-----------------------------")*/
	f, err := os.Create("ozon/test.txt")
	if loggerFile {
		if err != nil {
			panic(err)
		}
	}
	for idxTest, dataPost := range mapTestPost {
		if logger {
			fmt.Printf("Тест номер %d\n", idxTest+1)
		}
		//	fmt.Println(dataPost.comments)
		idCommet := make([]int, 0, dataPost.count)
		for id, _ := range dataPost.comments {
			idCommet = append(idCommet, id)
		}
		sort.Ints(idCommet)

		sort.SliceStable(idCommet, func(i, j int) bool {
			if dataPost.comments[idCommet[i]].idParent == dataPost.comments[idCommet[j]].idParent {
				if idCommet[i] < idCommet[j] {
					return true
				} else {
					return false
				}
			} else if dataPost.comments[idCommet[i]].idParent < dataPost.comments[idCommet[j]].idParent {
				return true
			} else {
				return false
			}
		})

		res := make([]Comment, 0, dataPost.count)
		findChild(-1, dataPost.comments, idCommet, 0, &res)

		resText := make([]string, 0, len(res))
		shift := "   "
		for idx := 0; idx < len(res); idx++ {
			comment := res[idx]
			s := comment.text
			if comment.level > 1 {
				s = "|--" + s
			}
			level := ""
			if comment.level > 1 {
				for i := 1; i < comment.level-1; i++ {
					s = shift + s
					level = shift + level
				}
			}

			if comment.level == 1 && len(resText) > 0 {
				resText = append(resText, "")
			}
			if comment.level > 1 {
				resText = append(resText, level+"|")
				idxSymbol := len(level)
				for i := len(resText) - 1; i > 0; i-- {
					if len(resText[i]) > idxSymbol {
						if resText[i-1] != "" && resText[i][idxSymbol] == ' ' && (idxSymbol-1 <= 0 || (idxSymbol-1 > 0 && resText[i][idxSymbol-1] == ' ')) {
							resText[i] = resText[i][0:idxSymbol] + "|" + resText[i][idxSymbol+1:]
						}
					} else {
						break
					}
				}
			}
			resText = append(resText, s)
		}

		for _, s := range resText {
			if loggerFile {
				_, err = f.WriteString(s + "\n")
				if err != nil {
					panic(err)
				}
			}

			fmt.Println(s)
		}
		fmt.Println()
		if loggerFile {
			_, err = f.WriteString("\n")
			if err != nil {
				panic(err)
			}
		}
	}
	if loggerFile {
		defer f.Close()
	}

}

func findChild(idParent int, comments map[int]Comment, idCommet []int, level int, res *[]Comment) (int, bool) {
	var ok bool
	for _, id := range idCommet {
		comment := comments[id]
		if comment.use {
			continue
		}
		if comment.idParent == idParent {
			level++
			comment.use = true
			comment.level = level
			*res = append(*res, comment)
			level, ok = findChild(comment.id, comments, idCommet, level, res)
			if !ok {
				level--
			}
		}
	}
	return level, ok
}
