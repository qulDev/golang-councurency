package helper

import "testing"

// var name []string = []string{"Alice", "Bob", "Charlie"}

// Data untuk testing
var testData = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var testUrls = []string{
	"https://example1.com/file1.zip",
	"https://example2.com/file2.zip",
	"https://example3.com/file3.zip",
	"https://example4.com/file4.zip",
	"https://example5.com/file5.zip",
}
var largeNumbers = make([]int, 1000)

func init() {
	// Inisialisasi data untuk benchmark matematika
	for i := 0; i < 1000; i++ {
		largeNumbers[i] = i + 1
	}
}

// ========== BENCHMARK GREETING FUNCTIONS ==========

// // BenchmarkGreetAll-12               10000            113253 ns/op             114 B/op
// func BenchmarkGreetAll(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GreetAllWithGoroutine(name)
// 	}
// }

// // BenchmarkGreetAllWithoutGoroutine-12               10000            111613 ns/op               0 B/op
// func BenchmarkGreetAllWithoutGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GreetAll(name)
// 	}
// }

// ========== BENCHMARK DATA PROCESSING ==========

func BenchmarkProcessDataSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataSequential(testData)
	}
}

func BenchmarkProcessDataConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataConcurrent(testData)
	}
}

// ========== BENCHMARK DOWNLOAD SIMULATION ==========

func BenchmarkDownloadFilesSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DownloadFilesSequential(testUrls)
	}
}

func BenchmarkDownloadFilesConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DownloadFilesConcurrent(testUrls)
	}
}

// ========== BENCHMARK MATHEMATICAL COMPUTATION ==========

func BenchmarkCalculateSumSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSumSequential(largeNumbers)
	}
}

func BenchmarkCalculateSumConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSumConcurrent(largeNumbers)
	}
}

// ========== UNIT TESTS ==========

func TestProcessDataSequential(t *testing.T) {
	input := []int{2, 3, 4}
	expected := []int{4, 9, 16}
	result := ProcessDataSequential(input)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d at index %d", expected[i], v, i)
		}
	}
}

func TestProcessDataConcurrent(t *testing.T) {
	input := []int{2, 3, 4}
	expected := []int{4, 9, 16}
	result := ProcessDataConcurrent(input)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d at index %d", expected[i], v, i)
		}
	}
}

func TestCalculateSumSequential(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := CalculateSumSequential(input)
	expected := 15 // 1+2+3+4+5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCalculateSumConcurrent(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := CalculateSumConcurrent(input)
	expected := 15 // 1+2+3+4+5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
