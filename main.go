package main

import (
	"runtime"
	"sync/atomic"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	
	fmt.Println("Cpus: \t", runtime.NumCPU())
	goruntimeCount := 10
	var contador int64
	contador = 0

	wg.Add(goruntimeCount)

	for i := 0; i < goruntimeCount; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println("Contador: \t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final: ", contador)
	
}