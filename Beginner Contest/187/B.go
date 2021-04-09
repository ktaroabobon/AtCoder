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

func findTilt(points [][]int) (r int) {
	r = 0

	for i := 0; i <= len(points)-2; i++ {
		for j := i + 1; j <= len(points)-1; j++ {
			t := (float64(points[j][1]) - float64(points[i][1])) / (float64(points[j][0]) - float64(points[i][0]))
			if -1 <= t && t <= 1 {
				r++
			}
		}
	}
	return
}

func main() {
	n := num()
	//n := 10
	if n == 1 {
		fmt.Println(0)
		return
	}

	points := make([][]int, 0)
	//points := [][]int{
	//	[]int{-31, -35},
	//	[]int{8, -36},
	//	[]int{22, 64},
	//	[]int{5, 73},
	//	[]int{-14, 8},
	//	[]int{18, -58},
	//	[]int{-41, -85},
	//	[]int{1, -88},
	//	[]int{-21, -85},
	//	[]int{-11, 82},
	//}

	for i := 0; i < n; i++ {
		ns := nums()
		points = append(points, ns)
	}

	r := findTilt(points)
	fmt.Println(r)
}
