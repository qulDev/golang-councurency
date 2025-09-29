package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	fmt.Println("Hello from goroutine", index)
	time.Sleep(1 * time.Second)
}
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RunAsync(&wg, i)
	}
	wg.Wait()
	fmt.Println("All goroutines complete")

}
