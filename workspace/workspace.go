package main

import "fmt"

// page URL:

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {

	type Stringer interface {
		String() string
	}
	var s Stringer = Hex(100)
	fmt.Println(s.String())
}
