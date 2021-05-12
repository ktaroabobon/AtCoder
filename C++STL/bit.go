package main

import "fmt"

const (
	A uint = 10 // 1010
	B uint = 12 // 1100
)

func main() {
	var bits uint

	fmt.Println(A)
	fmt.Println(B)

	// AND
	bits = A & B // 1001
	fmt.Println(bits)

	// OR
	bits = A | B // 1110
	fmt.Println(bits)

	// XOR
	bits = A ^ B // 0110
	fmt.Println(bits)

	// NAND
	bits = A &^ B // 0010
	fmt.Println(bits)

	//	左シフト
	bits = 1 << 3 // 1000
	fmt.Println(bits)

	// 右シフト
	bits = 8 >> 3 // 0001
	fmt.Println(bits)

	bits = 15 >> 2 // 1111 -> 0011 右に動かす
	fmt.Println(bits)

	fmt.Println("========================================")
	n := 2

	for bs := 0; bs < (1 << uint64(n)); bs++ {
		//　n桁分のbit列をforループで回す
		for i := 0; i < n; i++ {
			// 設定したbit列に関して一桁ずつ確認する
			fmt.Println(bs>>uint64(i)&1 == 1)
			// 指定した桁(i)は1であるか
		}
	}
}
