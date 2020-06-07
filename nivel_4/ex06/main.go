package main

import "fmt"

func main() {
	
	x := make([]string, 3, 3)
	estados := []string{"ce", "pe", "pb"}

	fmt.Println("len: ", len(estados))

	fmt.Println("cap", cap(estados))

	for i := 0; i < len(x); i++ {
		fmt.Println(estados[i])
	}
}