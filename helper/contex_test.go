package helper

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stop:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()

	time.Sleep(600 * time.Millisecond)
	cancel()
	time.Sleep(600 * time.Millisecond)
}

func TestContextWithTImeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stop:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()

	time.Sleep(600 * time.Millisecond)
	cancel()
	time.Sleep(600 * time.Millisecond)
}
