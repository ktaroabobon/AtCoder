package main

import (
	"bufio"
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
文字列をSplitしてSliceを返す
e.g.)
foo boo -> [foo boo]
*/
func splitWithoutEmpty(strTargeted string) (strReturned []string) {
	strSplit := strings.Split(strTargeted, " ")

	for _, v := range strSplit {
		if v != "" {
			strReturned = append(strReturned, v)
		}
	}

	return

}

func splitToInt(strTargeted string) (intReturned []int, err error) {
	strSplit := splitWithoutEmpty(strTargeted)

	for _, str := range strSplit {
		var v int
		v, err = strconv.Atoi(str)
		if err != nil {
			return
		}

		intReturned = append(intReturned, v)
	}

	return
}

/*
文字列型、１行読み込み
e.g.)
foo boo
*/
func strs() (strSlice []string) {
	str := condScan()

	strSlice = splitWithoutEmpty(str)

	return
}

/*
数値型、１行読み込み
e.g.)
100 200
*/
func nums() (intSlice []int) {
	str := condScan()

	intSlice, _ = splitToInt(str)

	return
}

func main() {

}
