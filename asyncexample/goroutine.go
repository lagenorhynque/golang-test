package asyncexample

import (
	"fmt"
	"math/rand"
	"time"
)

func TestGoroutine() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			sleep := time.Duration(rand.Intn(1000))
			time.Sleep(sleep * time.Millisecond)
			fmt.Println(i)
		}(i)
	}

	time.Sleep(2000 * time.Millisecond)
}
