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

func isOptionalValues(v, a, b int) (f bool) {
	sum := 0
	for {
		sum += (v % 10)
		v /= 10
		if v == 0 {
			break
		}
	}

	if a <= sum && sum <= b {
		f = true
	} else {
		f = false
	}

	return
}

func main() {
	nums := nums()
	n, a, b := nums[0], nums[1], nums[2]
	//n, a, b := 100, 4, 16
	r := 0

	for i := 0; i <= n; i++ {
		if isOptionalValues(i, a, b) {
			r += i
		}
	}

	fmt.Println(r)

}
