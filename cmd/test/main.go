package main

import (
	"fmt"
	"time"
)

// просто для теста
func main() {

	render(doubler(writer()))
}

func render(r chan int) {
	for v := range r {
		fmt.Println(v)
	}
}

func doubler(c chan int) chan int {

	var d = make(chan int)

	go func(c chan int, d chan int) {
		for v := range c {
			d <- v * 2
			time.Sleep(time.Millisecond * 500)
		}
		close(d)
	}(c, d)

	return d

}

func writer() chan int {
	var c = make(chan int)
	go func(c chan int) {
		for v := range 10 {
			c <- v
		}
		close(c)
	}(c)

	return c
}
