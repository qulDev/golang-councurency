package helper

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	//var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//once.Do(OnlyOnce)
			OnlyOnce()
		}()

	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
