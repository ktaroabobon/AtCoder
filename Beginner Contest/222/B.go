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

// page URL: https://atcoder.jp/contests/abc222/tasks/abc222_b

var N, P, cnt int
var AS []int

/*
main関数
*/

func main() {
	is := isReader()
	N, P = is[0], is[1]
	AS = isReader()

	for i := 0; i < N; i++ {
		a := AS[i]
		if a < P {
			cnt++
		}
	}

	fmt.Println(cnt)

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

func iswap(v1, v2 *int) {
	*v1, *v2 = *v2, *v1
}

func imodinv(x, y int) int {
	/*
		mod y における逆元を求めるアルゴリズム
		<逆元の存在条件>
		mod p でのaの逆元が存在する条件は、pとaとが互いに素であること
	*/

	z, u, v := y, 1, 0
	for {
		if !i2b(z) {
			break
		}
		t := x / z
		x -= t * z
		iswap(&x, &z)
		u -= t * v
		iswap(&u, &v)
	}
	u %= y
	if u < 0 {
		u += y
	}
	return u
}

func ichmin(a *int, b int) (f bool) {
	if *a > b {
		*a = b
		f = true
	}
	return
}

func ichmax(a *int, b int) (f bool) {
	if *a < b {
		*a = b
		f = true
	}
	return
}
func aCb(a, b int) (r int) {
	r = 1
	if a < b*2 {
		b = a - b
	}
	for i := 1; i < b+1; i++ {
		r = r * (a - i + 1) / i
	}
	return
}

/*
その他関数
*/
/* strSlice内に対象の文字列が存在するか*/
func ssContain(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if s == v {
			return true
		}
	}
	return false
}

/*stringSliceを逆順にして返す。*/
func ssReverse(data []string) []string {
	if len(data) == 0 {
		return data
	}
	return append(ssReverse(data[1:]), data[0])
}

/*stringSliceの初期化*/
func initSS(ss []string, v string) []string {
	if cap(ss) == 0 {
		return ss
	}
	if len(ss) == 0 {
		for i := 0; i < cap(ss); i++ {
			ss = append(ss, v)
		}
	} else {
		for i := 0; i < cap(ss); i++ {
			ss[i] = v
		}
	}
	return ss
}

/* intSlice内に対象の数値が存在するか*/
func isContain(intSlice []int, i int) bool {
	for _, v := range intSlice {
		if i == v {
			return true
		}
	}
	return false
}

/*intSliceを逆順にして返す。*/
func isReverse(data []int) []int {
	if len(data) == 0 {
		return data
	}
	return append(isReverse(data[1:]), data[0])
}

/*intSliceの初期化*/
func initIS(is []int, v int) []int {
	if cap(is) == 0 {
		return is
	}
	if len(is) == 0 {
		for i := 0; i < cap(is); i++ {
			is = append(is, v)
		}
	} else {
		for i := 0; i < cap(is); i++ {
			is[i] = v
		}
	}
	return is
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

/* Union Find */
type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.par = make([]int, n)
	uf.rank = make([]int, n)
	uf.size = make([]int, n)
	for i := range uf.par {
		uf.par[i] = -1
		uf.rank[i] = 0
		uf.size[i] = 1
	}
	return uf
}

func (u UnionFind) InitSize(s []int) {
	u.size[0] = 0
	for i := 1; i < len(u.size); i++ {
		u.size[i] = s[i] - s[i-1]
	}
}

func (u UnionFind) Root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.Root(u.par[x])
	return u.par[x]
}

func (u UnionFind) Unite(x, y int) {
	xr := u.Root(x)
	yr := u.Root(y)

	if xr == yr {
		return
	}
	// rank
	if u.rank[xr] < u.rank[yr] {
		u.par[xr] = yr
		u.size[yr] += u.size[xr]
	} else {
		u.par[yr] = xr
		u.size[xr] += u.size[yr]
		if u.rank[xr] == u.rank[yr] {
			u.rank[xr]++
		}
	}
}

func (u UnionFind) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u UnionFind) Size(x int) int {
	return u.size[u.Root(x)]
}
