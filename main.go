package main

import (
	"fmt"
	"time"
)

// Simulasi operasi yang memakan waktu
func fetchDataFromAPI(id int) string {
	time.Sleep(2 * time.Second) // Simulasi network delay
	return fmt.Sprintf("Data dari API %d", id)
}

// SYNCHRONOUS - Satu per satu, blocking
func synchronousExample() {
	fmt.Println("=== SYNCHRONOUS EXAMPLE ===")
	start := time.Now()

	// Panggil satu per satu, harus menunggu
	data1 := fetchDataFromAPI(1)
	data2 := fetchDataFromAPI(2)
	data3 := fetchDataFromAPI(3)

	fmt.Printf("Data 1: %s\n", data1)
	fmt.Printf("Data 2: %s\n", data2)
	fmt.Printf("Data 3: %s\n", data3)

	fmt.Printf("Total waktu: %v\n", time.Since(start))
	// Akan memakan waktu ~6 detik (2+2+2)
}

// ASYNCHRONOUS - Menggunakan goroutine + channel
func asynchronousExample() {
	fmt.Println("\n=== ASYNCHRONOUS EXAMPLE ===")
	start := time.Now()

	// Buat channel untuk komunikasi
	ch := make(chan string, 3) // Buffer untuk 3 hasil

	// Jalankan 3 goroutine bersamaan
	go func() {
		result := fetchDataFromAPI(1)
		ch <- result // Kirim hasil ke channel
	}()

	go func() {
		result := fetchDataFromAPI(2)
		ch <- result
	}()

	go func() {
		result := fetchDataFromAPI(3)
		ch <- result
	}()

	// Terima hasil dari channel
	for i := 0; i < 3; i++ {
		data := <-ch // Ambil data dari channel
		fmt.Printf("Received: %s\n", data)
	}

	fmt.Printf("Total waktu: %v\n", time.Since(start))
	// Akan memakan waktu ~2 detik (karena parallel)
}

// Contoh dengan return value menggunakan channel
func fetchDataAsync(id int, ch chan<- string) {
	// Simulasi operasi async
	go func() {
		time.Sleep(1 * time.Second)
		result := fmt.Sprintf("Async data %d", id)
		ch <- result // Kirim hasil ke channel
	}()
}

func channelReturnExample() {
	fmt.Println("\n=== CHANNEL AS RETURN VALUE ===")

	// Buat channel
	resultCh := make(chan string)

	// Panggil function async
	fetchDataAsync(42, resultCh)

	// Tunggu dan ambil hasil
	fmt.Println("Menunggu hasil...")
	result := <-resultCh // Blocking sampai ada data
	fmt.Printf("Hasil: %s\n", result)
}

// Contoh dengan multiple channels dan select
func selectExample() {
	fmt.Println("\n=== SELECT EXAMPLE (Non-blocking) ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine yang cepat
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data cepat"
	}()

	// Goroutine yang lambat
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Data lambat"
	}()

	// Select: ambil yang pertama datang
	select {
	case msg1 := <-ch1:
		fmt.Printf("Dapat dari ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Dapat dari ch2: %s\n", msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Tidak ada data dalam 2 detik")
	}
}

func main() {
	// Jalankan contoh synchronous
	synchronousExample()

	// Jalankan contoh asynchronous
	asynchronousExample()

	// Contoh channel sebagai return value
	channelReturnExample()

	// Contoh select
	selectExample()
}
