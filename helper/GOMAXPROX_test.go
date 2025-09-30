package helper

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprox(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

}
