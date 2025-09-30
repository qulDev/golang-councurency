package helper

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(wg *sync.WaitGroup, data *sync.Map, value int) {
	defer wg.Done()

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	var data sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go AddToMap(&wg, &data, i)
	}

	wg.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
