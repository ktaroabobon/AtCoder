// 深さ優先探索

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

//入力情報
//URL: https://qiita.com/drken/items/4a7869c5e304883f539b#3-4-dfs-%E3%81%AE%E6%8E%A2%E7%B4%A2%E9%A0%86%E5%BA%8F%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%AE%E8%A9%B3%E7%B4%B0
//
//15 14
//0 1
//0 4
//0 11
//1 2
//1 3
//4 5
//4 8
//5 6
//5 7
//8 9
//8 10
//11 12
//11 13
//13 14

type Graph = [][]int
type seen = []bool

var N, M int
var depth, subTreeSz []int

func dfs(g Graph, s seen, v int) {
	//行きがけ順
	fmt.Println(v)

	s[v] = true

	for _, nextV := range g[v] {
		if s[nextV] {
			continue
		}
		dfs(g, s, nextV)
		//fmt.Println(v)
	}
}
func dfsLast(g Graph, s seen, v int) {
	//帰りがけ順
	s[v] = true

	for _, nextV := range g[v] {
		if s[nextV] {
			continue
		}
		dfsLast(g, s, nextV)
	}

	fmt.Println(v)
}
func dfsTS(g Graph, s seen, v, t int) int {
	//timestampバージョン
	fmt.Printf("No.%v %v\n", t, v)
	t++

	s[v] = true

	for _, nextV := range g[v] {
		if s[nextV] {
			continue
		}
		t = dfsTS(g, s, nextV, t)
	}

	fmt.Printf("No.%v %v\n", t, v)
	t++

	return t
}

func dfsOnTree(g Graph, v, p, d int) {
	//	URL: https://qiita.com/drken/items/a803d4fc4a727e02f7ba#4-4-%E6%A0%B9%E3%81%AE%E3%81%AA%E3%81%84%E6%9C%A8%E3%81%AE%E8%B5%B0%E6%9F%BB
	//	根のない木の走査
	/*
		   	与えられた木の頂点 0 を根として指定した根つき木において、
		      ・各頂点の根からの距離 (深さ (depth) といいます)
		      ・各頂点を根とした部分木のサイズ (部分木に含まれる頂点数)
			をそれぞれ求めてみましょう。
	*/

	// 行きがけ各頂点の深さを求める
	depth[v] = d
	for _, nv := range g[v] {
		if nv == p {
			continue
		}
		dfsOnTree(g, nv, v, d+1)
	}

	//	帰りがけに部分木のサイズを求める
	subTreeSz[v] = 1
	for _, c := range g[v] {
		if c == p {
			continue
		}
		subTreeSz[v] += subTreeSz[c]
	}
	/*
		総じて、
		・各頂点の深さ: 行きがけ順に求めた
		・各頂点を根とした部分木のサイズ: 帰りがけ順に求めた
		といえます。
	*/
}

/*
main関数
*/

func main() {
	numbers := isReader()
	N, M = numbers[0], numbers[1]

	g := make(Graph, N)

	for i := 0; i < M; i++ {
		info := isReader()
		g[info[0]] = append(g[info[0]], info[1])
		g[info[1]] = append(g[info[1]], info[0])
	}

	fmt.Println(g)

	s := make(seen, N)
	fmt.Println(s)

	fmt.Println("DFS 行きがけ順")
	dfs(g, s, 0)

	s = make(seen, N)
	//fmt.Println(s)

	fmt.Println("DFS 帰りがけ順")
	dfsLast(g, s, 0)

	s = make(seen, N)
	//fmt.Println(s)

	fmt.Println("DFS タイムスタンプ")
	dfsTS(g, s, 0, 0)

	fmt.Println("DFS 根のない木の走査")
	root := 0
	depth = make([]int, N)
	subTreeSz = make([]int, N)
	dfsOnTree(g, root, -1, 0)

	for i := 0; i < N; i++ {
		fmt.Printf("No.%v Depth:%v subTreeSz:%v\n", i, depth[i], subTreeSz[i])
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

// Deque
func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) Push(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Inject(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) Pop() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Eject() interface{} {
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
