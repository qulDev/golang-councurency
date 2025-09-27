package helper

import (
	"fmt"
	"sync"
	"time"
)

func SayHelloWithGoroutine(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Hello, " + name)
}

func GreetAllWithGoroutine(names []string) {
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go SayHelloWithGoroutine(name, &wg)
	}
	wg.Wait()
	print("All greetings sent.\n")
}

func SayHello(name string) {
	println("Hello, " + name)
}

func GreetAll(names []string) {
	for _, name := range names {
		SayHello(name)
	}
	print("All greetings sent.\n")
}

// ========== CONTOH KASUS: SIMULASI OPERASI BERAT ==========

// ProcessDataSequential - Memproses data secara berurutan (tanpa goroutine)
func ProcessDataSequential(data []int) []int {
	result := make([]int, len(data))

	for i, value := range data {
		// Simulasi operasi berat (misalnya: komputasi kompleks, I/O operation)
		time.Sleep(10 * time.Millisecond)
		result[i] = value * value // Menghitung kuadrat
	}

	return result
}

// ProcessDataConcurrent - Memproses data secara bersamaan (dengan goroutine)
func ProcessDataConcurrent(data []int) []int {
	result := make([]int, len(data))
	var wg sync.WaitGroup

	for i, value := range data {
		wg.Add(1)
		go func(index int, val int) {
			defer wg.Done()
			// Simulasi operasi berat yang sama
			time.Sleep(10 * time.Millisecond)
			result[index] = val * val
		}(i, value)
	}

	wg.Wait()
	return result
}

// ========== CONTOH KASUS: DOWNLOAD SIMULATION ==========

// DownloadFilesSequential - Simulasi download file secara berurutan
func DownloadFilesSequential(urls []string) []string {
	results := make([]string, len(urls))

	for i, url := range urls {
		// Simulasi waktu download (50ms per file)
		time.Sleep(50 * time.Millisecond)
		results[i] = fmt.Sprintf("Downloaded: %s", url)
	}

	return results
}

// DownloadFilesConcurrent - Simulasi download file secara bersamaan
func DownloadFilesConcurrent(urls []string) []string {
	results := make([]string, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(index int, fileUrl string) {
			defer wg.Done()
			// Simulasi waktu download yang sama
			time.Sleep(50 * time.Millisecond)
			results[index] = fmt.Sprintf("Downloaded: %s", fileUrl)
		}(i, url)
	}

	wg.Wait()
	return results
}

// ========== CONTOH KASUS: MATHEMATICAL COMPUTATION ==========

// CalculateSumSequential - Menghitung sum dari array besar secara berurutan
func CalculateSumSequential(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		// Simulasi operasi matematika yang lebih berat
		for j := 0; j < 1000; j++ {
			sum += num
		}
	}
	return sum / 1000 // Rata-rata dari hasil operasi
}

// CalculateSumConcurrent - Menghitung sum dengan membagi pekerjaan ke goroutines
func CalculateSumConcurrent(numbers []int) int {
	numWorkers := 4
	if len(numbers) < numWorkers {
		numWorkers = len(numbers)
	}

	if numWorkers == 0 {
		return 0
	}

	chunkSize := len(numbers) / numWorkers
	remainder := len(numbers) % numWorkers

	resultChan := make(chan int, numWorkers)
	var wg sync.WaitGroup

	start := 0
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		currentChunkSize := chunkSize
		if i < remainder {
			currentChunkSize++
		}

		go func(startIdx, endIdx int) {
			defer wg.Done()

			localSum := 0
			for j := startIdx; j < endIdx; j++ {
				// Simulasi operasi matematika yang sama
				for k := 0; k < 1000; k++ {
					localSum += numbers[j]
				}
			}
			resultChan <- localSum / 1000
		}(start, start+currentChunkSize)

		start += currentChunkSize
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalSum := 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	return totalSum
}
