package main

import (
	"fmt"
	"sync"
)

func sum(a int) int {
	return a * a
}

func out(fn func(int) int, in1, in2, out chan int, n int) {
	arr1 := make([]int, n)
	arr2 := make([]int, n)
	m := new(sync.Mutex)
	j := 0
	for i := 0; i < n; i++ {
		go func(m *sync.Mutex) {
			m.Lock()
			a := <-in1
			num := j
			b := <-in2
			j++
			m.Unlock()
			wg := new(sync.WaitGroup)
			wg.Add(2)
			go func(wg *sync.WaitGroup) {
				arr1[num] = fn(a)
				wg.Done()
			}(wg)
			go func(wg *sync.WaitGroup) {
				arr2[num] = fn(b)
				wg.Done()
			}(wg)
			go func(wg *sync.WaitGroup) {
				wg.Wait()
				out <- arr1[num] + arr2[num]
			}(wg)
		}(m)
	}
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	out(sum, c1, c2, c3, 10)
	for j := 0; j < 10; j++ {
		c1 <- j
		c2 <- j
		fmt.Println(<-c3)
	}
}
