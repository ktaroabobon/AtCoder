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
数値、１整数
e.g.)
foo
*/
func num() (numReturned int) {
	str := condScan()

	numReturned, _ = strconv.Atoi(strings.TrimSpace(str))

	return
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

func main() {
	n := num()
	if n == 1 {
		fmt.Println(1)
		return
	}
	sumAoki := 0
	all := make([]int, 0)

	for i := 0; i < n; i++ {
		v := nums()
		sumAoki += v[0]
		all = append(all, 2*v[0]+v[1])
	}

	sort.Slice(all, func(i, j int) bool { return all[i] >= all[j] })

	sumAll := 0
	for i, v := range all {
		sumAll += v
		if sumAll > sumAoki {
			fmt.Println(i + 1)
			return
		}
	}
}
