package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var contador int32 = 0

const quantidadeGoruntime = 5

func main() {
	criarGoroutine(quantidadeGoruntime)
	fmt.Println("Total gorutime: ", quantidadeGoruntime)
	wg.Wait()
}

func criarGoroutine(count int) {
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			fmt.Println(v)
			wg.Done()
		}()
	}
}
