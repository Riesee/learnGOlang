package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Merhaba %s! (%d)\n", name, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		sayHello("Ali")
	}()

	go func() {
		sayHello("Veli")
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("Program bitti")
}