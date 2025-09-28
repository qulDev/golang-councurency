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

func OnlyIn(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Rizqullah"

}

func OnlyOut(ch <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-ch
	fmt.Println(data)
}

func TestOnlyInOnlyOut(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	fmt.Println("Menunggu...")
	time.Sleep(3 * time.Second)

}
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Rizqullah"

	fmt.Println("selesai")
}
func TestRangeChannel(t *testing.T) {
	channel := make(chan string, 9)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Rizqullah %d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespond(channel1)
	go GiveMeRespond(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++

		}

		if counter == 2 {
			break
		}

	}
}
func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespond(channel1)
	go GiveMeRespond(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++

		default:
			fmt.Println("Menunggu data...")

		}

		if counter == 2 {
			break
		}

	}
}
