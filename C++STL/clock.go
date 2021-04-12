package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 5)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}
