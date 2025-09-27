# Go Concurrency Benchmark Examples

Proyek ini mendemonstrasikan perbedaan performa antara operasi **sequential** (berurutan) dan **concurrent** (bersamaan) menggunakan goroutines di Go.

## 📊 Hasil Benchmark

### 1. Data Processing (Operasi Komputasi)
```
BenchmarkProcessDataSequential-12     10    101,780,227 ns/op    92 B/op     1 allocs/op
BenchmarkProcessDataConcurrent-12    100     10,330,195 ns/op  2,215 B/op    32 allocs/op
```
**Hasil:** Concurrent **~10x lebih cepat** dari sequential! 🚀

### 2. Download Simulation (I/O Operations)
```
BenchmarkDownloadFilesSequential-12    4    250,994,380 ns/op    850 B/op    11 allocs/op
BenchmarkDownloadFilesConcurrent-12   22     50,307,348 ns/op  1,511 B/op    27 allocs/op
```
**Hasil:** Concurrent **~5x lebih cepat** dari sequential! 🚀

### 3. Mathematical Computation (CPU-bound)
```
BenchmarkCalculateSumSequential-12   8478       135,894 ns/op      0 B/op     0 allocs/op
BenchmarkCalculateSumConcurrent-12   7736       142,399 ns/op    514 B/op    11 allocs/op
```
**Hasil:** Sequential sedikit lebih cepat untuk operasi CPU-bound sederhana karena overhead goroutines.

## 🎯 Contoh Kasus yang Diimplementasikan

### 1. **Data Processing** - Pemrosesan Array
**Sequential Version:**
- Memproses setiap elemen array satu per satu
- Waktu total = n × waktu_per_operasi

**Concurrent Version:**
- Memproses semua elemen secara bersamaan dengan goroutines
- Waktu total ≈ waktu_per_operasi (terlama)

```go
// Sequential
func ProcessDataSequential(data []int) []int {
    result := make([]int, len(data))
    for i, value := range data {
        time.Sleep(10 * time.Millisecond) // Simulasi operasi berat
        result[i] = value * value
    }
    return result
}

// Concurrent
func ProcessDataConcurrent(data []int) []int {
    result := make([]int, len(data))
    var wg sync.WaitGroup
    
    for i, value := range data {
        wg.Add(1)
        go func(index int, val int) {
            defer wg.Done()
            time.Sleep(10 * time.Millisecond) // Operasi yang sama
            result[index] = val * val
        }(i, value)
    }
    wg.Wait()
    return result
}
```

### 2. **Download Simulation** - Operasi I/O
Mensimulasikan download file dari multiple URLs secara bersamaan vs berurutan.

### 3. **Mathematical Computation** - Komputasi Paralel
Membagi pekerjaan matematika berat ke multiple workers.

## 🏃‍♂️ Cara Menjalankan

### Menjalankan Benchmark
```bash
go test ./helper -run=^$ -bench=. -benchmem
```

### Menjalankan Unit Tests
```bash
go test ./helper -v
```

### Menjalankan Benchmark Spesifik
```bash
# Hanya benchmark data processing
go test ./helper -run=^$ -bench=BenchmarkProcessData

# Hanya benchmark download simulation  
go test ./helper -run=^$ -bench=BenchmarkDownloadFiles

# Hanya benchmark mathematical computation
go test ./helper -run=^$ -bench=BenchmarkCalculateSum
```

## 📈 Analisis Hasil

### Kapan Menggunakan Goroutines?

**✅ Gunakan Goroutines untuk:**
- **I/O Operations** (file, network, database)
- **Independent Tasks** yang bisa dijalankan paralel
- **Waiting Operations** (sleep, timeout)
- **Multiple API Calls**

**❌ Hindari Goroutines untuk:**
- **Simple CPU-bound tasks** pada single core
- **Sequential dependencies** (task B butuh hasil task A)
- **Shared state** yang banyak membutuhkan synchronization

### Trade-offs

**Concurrent Advantages:**
- Performa lebih cepat untuk I/O operations
- Better resource utilization
- Improved user experience (non-blocking)

**Concurrent Costs:**
- Memory overhead (goroutine stack ~2KB)
- Context switching overhead
- Complexity in debugging
- Potential race conditions

## 🛠️ Struktur Proyek

```
golang-concurrency/
├── go.mod
├── main.go
├── README.md
└── helper/
    ├── sayHello.go      # Implementasi functions
    └── sayHello_test.go # Benchmark & unit tests
```

## 🎓 Pelajaran Penting

1. **Goroutines bukan silver bullet** - gunakan sesuai kebutuhan
2. **Measure, don't guess** - selalu benchmark untuk memastikan
3. **I/O bound tasks** mendapat benefit terbesar dari concurrency
4. **CPU bound tasks** perlu pertimbangan lebih hati-hati
5. **Memory usage** meningkat dengan penggunaan goroutines

Selamat belajar Go Concurrency! 🎉