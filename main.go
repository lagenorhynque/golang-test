package main

import (
	"fmt"

	"github.com/lagenorhynque/stringutil"
)

func fibonacci() func() int {
	prev, curr := 1, 0
	return func() int {
		prev, curr = curr, prev+curr
		return prev
	}
}

func main() {
	fmt.Printf("Hello, world!\n")
	fmt.Println(stringutil.Reverse("!dlrow ,olleH"))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
