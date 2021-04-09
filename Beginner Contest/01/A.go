package main

import (
	"bufio"
	"fmt"
	"os"
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
数値、１整数
e.g.)
foo
*/
func num() (numReturned int) {
	str := condScan()

	numReturned, _ = strconv.Atoi(strings.TrimSpace(str))

	return
}

func main() {
	h1, h2 := num(), num()

	fmt.Println(h1 - h2)
}
