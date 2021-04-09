package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

/*
数値、１整数
e.g.)
foo
*/
func num() (numReturned int) {
	str := condScan()

	numReturned, _ = strconv.Atoi(strings.TrimSpace(str))

	return
}

func isContain(words []string, s string) bool {
	for _, v := range words {
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	const ok = "satisfiable"

	n := num()
	if n == 1 {
		fmt.Println(ok)
		return
	}

	words := make([]string, 0)
	sMap := make(map[string]bool, 0)

	for i := 0; i < n; i++ {
		s := str()
		words = append(words, s)
	}

	sort.Strings(words)
	var sStart int
	for i, s := range words {
		if strings.Contains(s, "!") {
			sMap[s] = true
			continue
		}
		sStart = i
		break
	}

	words = words[sStart:]

	for _, v := range words {
		if sMap["!"+v] {
			fmt.Println(v)
			return
		}
	}

	fmt.Println(ok)

}
