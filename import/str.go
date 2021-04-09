package main

import (
	"bufio"
	"os"
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
