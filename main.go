package main

import (
	"fmt"
	"sync"
)

func sum(a int) int {
	return a * a
}
func change(in1, in2 chan int, n int) {
	var wg sync.WaitGroup
	var a, b int
	for i := 0; i < n; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			a = <-in1
			wg.Done()
		}(&wg)
		go func(wg *sync.WaitGroup) {
			b = <-in2
			wg.Done()
		}(&wg)
		wg.Wait()
		fmt.Println(a + b)
	}
}
func out(fn func(int) int, in1, in2, out chan int, n int) {
	for i := 0; i < n; i++ {
		go func() {
			var wg sync.WaitGroup

			out <- fn(<-in1) + fn(<-in2)
		}()
	}
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	/*for i:=0; i<10; i++ {
		go func() {
			//fmt.Println(<-c1)
			//fmt.Println(<-c2)
			c3 <- sum(<-c1) + sum(<-c2)
		}()
	}*/
	out(sum, c1, c2, c3, 10)
	for j := 0; j < 10; j++ {
		c1 <- j
		c2 <- j
		fmt.Println(<-c3)
	}
	//c2 <- 5
	/*for i := 0; i < 10; i++ {
		c1 <- i
		c2 <- i
	}*/
}
