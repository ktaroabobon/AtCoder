package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
float型で受け取る
*/
func num() (numReturned float64) {
	str := condScan()

	numReturned, _ = strconv.ParseFloat(str, 64)

	return
}

func main() {
	num := num() / 1000

	if num < 0.1 {
		fmt.Println("00")
	} else if 0.1 <= num && num <= 5 {
		fmt.Sprintf("%02d", num*100)
	} else if 6 <= num && num <= 30 {
		fmt.Println(int(num + 50))
	} else if 30 <= num && num <= 70 {
		fmt.Println(int((num-30)/5 + 80))
	} else {
		fmt.Println(89)
	}
}
