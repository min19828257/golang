package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int64
	mu sync.Mutex
}

// counter 값을 1씩 증가시킴
func (c *counter) increment() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

//counter 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	c.display()
}
