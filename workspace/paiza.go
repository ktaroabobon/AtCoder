package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
main関数
*/

func make_socks_map(n int) map[string]int {
	socks_map := make(map[string]int)
	for i := 0; i < n; i++ {
		num := 0
		ss := ssReader()
		kind, right_left := ss[0], ss[1]
		socks := kind + "_" + right_left

		if val, ok := socks_map[socks]; ok {
			num = val
		}

		socks_map[socks] = num + 1
	}
	return socks_map
}

func socks_count(socks_map map[string]int) (count int) {
	for key, val := range socks_map {
		kind := string(key[:len(key)-1])
		var new_key string

		if string(key[len(key)-1]) == "R" {
			new_key = kind + "L"
		} else {
			new_key = kind + "R"
		}

		num, _ := socks_map[new_key]

		if val < num {
			count += val
		} else {
			count += num
		}

	}

	return
}

func main() {
	n := iReader()

	socks_map := make_socks_map(n)

	count := socks_count(socks_map)

	fmt.Println(count / 2)
}

/*
標準入力の読み込み
*/
const BUFSIZE = 10000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

/*
スキャニング
*/
func readLine() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
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
		intReturned = append(intReturned, s2i(str))
	}

	return
}

/*
文字列、１単語
e.g.)
foo
*/
func sReader() (strReturned string) {
	strReturned = readLine()

	return
}

/*
数値、１整数
e.g.)
100
*/
func iReader() (numReturned int) {
	str := readLine()

	numReturned = s2i(str)

	return
}

/*
文字列型、１行読み込み
e.g.)
foo boo
*/
func ssReader() (strSlice []string) {
	str := readLine()

	strSlice = splitWithoutEmpty(str)

	return
}

/*
数値型、１行読み込み
e.g.)
100 200
*/
func isReader() (intSlice []int) {
	str := readLine()

	intSlice, _ = splitToInt(str)

	return
}

/*
型変換
*/
// string <-> int
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

func i2s(i int) string {
	return strconv.Itoa(i)
}

// bool <-> int
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func i2b(i int) bool {
	return i != 0
}

/*
int型計算式
*/
func iabs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func imin(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func imax(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func imod(x, y int) int {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
}

func ipow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func isum(values ...int) int {
	ret := 0
	for _, v := range values {
		ret += v
	}
	return ret
}

func igcd(v1, v2 int) int {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	for v1 != 0 {
		v1, v2 = v2%v1, v1
	}
	return v2
}

func ilcm(v1, v2 int) int {
	return v1 * v2 / igcd(v1, v2)
}

/*
その他関数
*/
/* strSlice内に対象の文字列が存在するか*/
func isContain(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if s == v {
			return true
		}
	}
	return false
}

/* intSlice内に対象の数値が存在するか*/
func iisContain(intSlice []int, i int) bool {
	for _, v := range intSlice {
		if i == v {
			return true
		}
	}
	return false
}
