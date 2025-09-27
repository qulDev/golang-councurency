package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello", name, "ke", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sayHello("Goroutine 1", &wg)
	go sayHello("Goroutine 2", &wg)
	go sayHello("Goroutine 3", &wg)
	wg.Wait()

	fmt.Println("Main function finished")
}
