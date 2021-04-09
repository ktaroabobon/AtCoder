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

func isSquere(intTargeted int) {
	for i := 2; i*i <= intTargeted; i++ {
		if i*i == intTargeted {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")

}

func main() {
	strSlice := strs()

	s := strings.Join(strSlice, "")
	n, _ := strconv.Atoi(s)

	isSquere(n)

}
