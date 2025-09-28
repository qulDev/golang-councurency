package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	ch := make(chan string)

	ch <- "Test Channel"

	close(ch)
}

func GiveMeRespond(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Rizqullah"
	fmt.Println("Selesai")
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go GiveMeRespond(ch)
	fmt.Println("Menunggu...")
	res := <-ch
	fmt.Println(res)

	time.Sleep(3 * time.Second)
}
