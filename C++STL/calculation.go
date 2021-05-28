package main

import "fmt"

/*
大きな数の計算
*/

var MOD int

func main() {
	MOD = 1000000007

	call()

	//fmt.Println(-17 % 5)

	a := 2000000020
	b := 20

	fmt.Printf("普通に計算して余りを求める: %v\n", (a-b)%MOD)
	fmt.Printf("余り求めてから計算して余りを求める: %v\n", ((a%MOD)-(b%MOD))%MOD)
	fmt.Printf("余り求めてから計算して余りを求める (対策済):%v\n", imod((a%MOD)-(b%MOD), MOD))

	a = 12345678900000
	b = 100000

	a %= MOD
	fmt.Println(a * imodinv(b, MOD) % MOD)
}

func call() {
	a := 111111111
	b := 123456789
	c := 987654321

	fmt.Println((a * b * c) % MOD)
	// 計算途中でオーバーフローしている
	fmt.Println(((a * b) % MOD * c) % MOD)
}

func imod(x, y int) int {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
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

func i2b(i int) bool {
	return i != 0
}
