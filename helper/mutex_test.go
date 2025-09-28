package helper

import (
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x int = 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Logf("Counter: %d", x)
}
