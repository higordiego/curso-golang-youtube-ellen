package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		fmt.Println("Primeira goruntime")
		wg.Done()
	}()

	go func() {
		fmt.Println("Segunda goruntime")
		wg.Done()
	}()

	wg.Wait()
}
