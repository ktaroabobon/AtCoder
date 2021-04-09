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
func isIn(v int, nums []int) bool {
	for _, n := range nums {
		if v == n {
			return true
		}
	}
	return false
}

func main() {
	n := num()
	nums := make([]int, 0)

	for i := 0; i < n; i++ {
		v := num()
		if !isIn(v, nums) {
			nums = append(nums, v)
		}

	}

	fmt.Println(len(nums))
}
