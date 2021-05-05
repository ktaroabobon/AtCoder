package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// page URL: https://atcoder.jp/contests/atc001/tasks/dfs_a

// DFSの問題です。

type Graph = [][]string

type seen = []bool

func initValueSlice(s [][]bool, v int) [][]bool {
	if cap(s) == 0 {
		return s
	}
	//x := make([]bool, v)
	for i := 0; i < cap(s); i++ {
		s = append(s, make(seen, v))
	}

	//for i := 1; i < cap(s); i *= 2 {
	//	copy(s[i:], s[:i])
	//}
	//s0 := s[0][0]
	//s1 := s[1][0]
	//s3 := s[1][1]
	//fmt.Printf("%v\n", &s[0][0])
	//fmt.Printf("%v\n", &s[0][1])
	//fmt.Printf("%v\n", &s[1][0])
	//fmt.Printf("%v\n", &s[1][1])
	return s

}
func solve(g Graph, s [][]bool, dx, dy []int, H, W, h, w int) [][]bool {
	s[h][w] = true

	for i := 0; i < 4; i++ {
		nextH := h + dx[i]
		nextW := w + dy[i]

		if nextH < 0 || nextH >= H || nextW < 0 || nextW >= W {
			continue
		}

		if g[nextH][nextW] == "#" {
			continue
		}

		if s[nextH][nextW] {
			continue
		}

		s = solve(g, s, dx, dy, H, W, nextH, nextW)
	}
	return s
}

/*
main関数
*/

func main() {
	numbers := isReader()
	H, W := numbers[0], numbers[1]

	field := make(Graph, H)

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for i := 0; i < H; i++ {
		info := sReader()
		field[i] = strings.Split(info, "")
	}
	// o(H)

	var sh, sw, gh, gw int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if field[i][j] == "s" {
				sh = i
				sw = j
			}
			if field[i][j] == "g" {
				gh = i
				gw = j
			}
		}
	}
	// o(H*W)

	s := make([][]bool, 0, H)

	s = initValueSlice(s, W)
	//	 o(logH)

	r := solve(field, s, dx, dy, H, W, sh, sw)

	if r[gh][gw] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
数値型のSliceを文字列型のSliceに変換して返す
e.g.)
[100 200] -> ["100" "200"]
*/
func splitToString(intTargeted []int) (strReturned []string, err error) {
	for _, v := range intTargeted {
		strReturned = append(strReturned, i2s(v))
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
foo
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

func s2r(s string) rune {
	var r int32
	for _, v := range s {
		r = v
	}
	return r
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

func iswap(v1, v2 int) (int, int) {
	return v2, v1
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

func toReverse(data []interface{}) []interface{} {
	if len(data) == 0 {
		return data
	}
	return append(toReverse(data[1:]), data[0])
}

/* intHeap(優先度付きキュー) */
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] } // 昇順
/* func (h intHeap) Less(i, j int) bool { return h[i] > h[j] } // 降順 */
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 二分探索
// 数値型スライスのなかで対象の数値以上の初めて登場するインデックスを返す
func lowerBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] >= x })
	return
}

func designatedLowerBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = lowerBound(intTarget, x)
	if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		returnIndex = -1
	}
	return
}

// 数値型スライスのなかで対象の数値以上の最後に登場するインデックスを返す
func upperBound(intTarget []int, x int) (returnIndex int) {
	returnIndex = sort.Search(len(intTarget), func(i int) bool { return intTarget[i] > x }) - 1
	return
}

func designatedUpperBound(intTarget []int, x int) (returnIndex int, f bool) {
	returnIndex = upperBound(intTarget, x)
	if returnIndex == -1 {
		f = false
	} else if returnIndex < len(intTarget) && intTarget[returnIndex] == x {
		f = true
	} else {
		f = false
		returnIndex = -1
	}
	return
}

func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) AppendLeft(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Append(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) PopLeft() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Pop() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}
