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
文字列型のSliceを数値型のSliceに変換して返す
e.g.)
["100" "200"] -> [100 200]
*/
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
数値型、１行読み込み
e.g.)
100 200
*/
func nums() (intSlice []int) {
	str := condScan()

	intSlice, _ = splitToInt(str)

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

func main() {
	_ = num()
	cards := nums()

	sum1, sum2 := 0, 0

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] >= cards[j]
	})

	for i, v := range cards {
		if i%2 == 0 {
			sum1 += v
		} else {
			sum2 += v
		}
	}

	fmt.Println(sum1 - sum2)
}
