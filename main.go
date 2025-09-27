package main

import "sync"

func SayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Hello, " + name)
}

func main() {
	var wg sync.WaitGroup
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		wg.Add(1)
		go SayHello(name, &wg)
	}
	wg.Wait()
	print("All greetings sent.\n")
}
