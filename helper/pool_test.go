package helper

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	var wg sync.WaitGroup
	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get()
			log.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()

	}
	wg.Wait()
	log.Println("selesai")
}
