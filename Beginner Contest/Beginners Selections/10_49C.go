package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* 標準入力を読み込む */
var scanner = bufio.NewScanner(os.Stdin)

/*
スキャニング
*/
func condScan() string {
	scanner.Scan()
	return scanner.Text()
}

/*
文字列、一単語
e.g.)
foo
*/
func str() (strReturned string) {
	strReturned = condScan()

	return
}

func main() {
	s := str()

	//s = strings.Replace(s, "dream", "D", -1)
	//s = strings.Replace(s, "erase", "E", -1)
	//s = strings.Replace(s, "Der", "", -1)
	//s = strings.Replace(s, "Er", "", -1)
	//s = strings.Replace(s, "D", "", -1)
	//s = strings.Replace(s, "E", "", -1)
	//s = strings.TrimSpace(s)
	//
	//if s == "" {
	//	fmt.Println("YES")
	//} else {
	//	fmt.Println("NO")
	//}

	words := []string{"dream", "dreamer", "erase", "eraser"}
	T := ""
	for true {
		temp := T

		if s == T {
			fmt.Println("YES")
			return
		}
		for _, v := range words {
			if strings.HasSuffix(s, v+T) {
				T = v + T
			}
		}
		if temp == T {
			break
		}
	}
	fmt.Println("NO")
}
