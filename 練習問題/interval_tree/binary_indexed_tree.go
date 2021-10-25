package main

import (
	"fmt"
	"github.com/gevg/bit"
)

// URL: https://qiita.com/DaikiSuyama/items/7295f5160a51684554a7

type BIT []int

func NewBIT(n int) BIT {
	tree := make([]int, n)
	return tree
}
func (a BIT) Len() int { return len(a) }

func (a BIT) Add(i, x int) {
	l := a.Len()
	for k := i; k < l; k += k & -k {
		a[k] += x
	}
}

func (a BIT) Sum(i, j int) int {
	return a.SumSub(j) - a.SumSub(i-1)
}

func (a BIT) SumSub(i int) int {
	var s int
	if i == -1 {
		return s
	}
	for k := i; k > 0; k -= k & -k {
		s += a[k]
	}
	return s
}

func (a BIT) LowerBound(x int) int {
	if x < 0 {
		return 0
	} else {
		l := a.Len()
		i := 0
		r := 1

		for {
			if r >= l {
				break
			}
			for ln := r; ln > 0; ln = ln >> 1 {
				if i+ln < l && a[i+l] < x {
					x -= a[i+ln]
					i += ln
				}
			}
			r = r << 1
		}
		return i
	}
}

func main() {
	numbers := []int32{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	tree := bit.From(numbers)

	for i := range numbers {
		fmt.Printf("tree(%d) = %d\n", i, tree.Sum(i))
	}

	tree = bit.New(10)
	fmt.Println(tree)
	tree.Add(0, 1)
	tree.Set(1, 3)
	tree.Set(2, 5)

	//for i := range numbers {
	//	fmt.Printf("tree(%d) = %d\n", i, tree.Sum(i))
	//}

	tree.RangeAdd(3, []int32{7, 9, 11})
	tree.RangeSet(6, []int32{13, 15, 17, 19})

	//for i := range numbers {
	//	fmt.Printf("tree(%d) = %d\n", i, tree.Sum(i))
	//}

	sums := tree.RangeSum(2, 4)
	fmt.Println(sums)

	tree = bit.New()
	fmt.Println(tree)

	tree = bit.Append(tree, 1)
	tree = bit.Append(tree, []int32{3, 5, 7, 9}...)
	tree = bit.Append(tree, 11, 13)
	fmt.Println(tree)

	tree.Scale(5)
	fmt.Println(tree)

	tree.RangeScale(2, 8, 6)
	fmt.Println(tree)

	tree.Mul(5, -2)
	fmt.Println(tree)

	t := NewBIT(10)
	fmt.Println(t)

	t.Add(0, 1)
	t.Add(1, 3)
	t.Add(2, 5)
	fmt.Println(t)
}
