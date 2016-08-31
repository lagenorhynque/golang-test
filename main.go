package main

import (
	"fmt"

	"github.com/lagenorhynque/golang-test/asyncexample"
	"github.com/lagenorhynque/golang-test/stringutil"
)

func fibonacci() func() int {
	prev, curr := 1, 0
	return func() int {
		prev, curr = curr, prev+curr
		return prev
	}
}

func main() {
	fmt.Println("--- Hello world ---")
	fmt.Printf("Hello, world!\n")
	fmt.Println(stringutil.Reverse("!dlrow ,olleH"))

	fmt.Println("--- fibonacci ---")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("--- async example ---")
	fmt.Println("goroutine:")
	asyncexample.TestGoroutine()
	fmt.Println("channel:")
	asyncexample.TestChannel()
	fmt.Println("select:")
	asyncexample.TestSelect()
}
