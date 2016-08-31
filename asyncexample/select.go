package asyncexample

import (
	"fmt"
	"math/rand"
	"time"
)

func TestSelect() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := make(chan string)

	go func() {
		sleep := time.Duration(rand.Intn(1000))
		time.Sleep(sleep * time.Millisecond)
		c <- "Success!"
	}()

	select {
	case <-c:
		fmt.Println("Got a value!")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout!")
	}
}
