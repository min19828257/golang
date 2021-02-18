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
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	// 모든 고루틴이 완료될 때까지 대기
	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display()
}
