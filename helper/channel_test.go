package helper

import "testing"

func TestCreateChannel(t *testing.T) {
	ch := make(chan string)

	ch <- "Test Channel"

	close(ch)
}
