package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

var contador = 0

const quantidadeGoruntime = 5000

func main() {
	criarGoroutine(quantidadeGoruntime)
	wg.Wait()
	fmt.Println("Total gorutime: ", quantidadeGoruntime)
	fmt.Println("Contador: ", contador)
}

func criarGoroutine(count int) {
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
