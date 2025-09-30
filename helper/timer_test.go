package helper

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("After Func", time.Now())
		wg.Done()
	})
	fmt.Println("Before", time.Now())
	wg.Wait()
}
