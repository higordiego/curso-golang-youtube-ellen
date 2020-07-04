package main

import (
	"runtime"
	"sync"
	"fmt"
)

var wg sync.WaitGroup

var contador = 0

const quantidadeGoruntime = 5000

func main() {
	criarGoroutine(quantidadeGoruntime)
	fmt.Println("Total gorutime: ", quantidadeGoruntime)
	wg.Wait()
}

func criarGoroutine(count int) {
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}
