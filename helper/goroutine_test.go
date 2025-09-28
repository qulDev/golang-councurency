package helper

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, World!")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)

}

func DisplayNumber(number int) {
	fmt.Println("Display: ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 1; i <= 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
