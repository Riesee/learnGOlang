package main

import (
	"fmt"
	"time"
)

// DERS 7: Goroutines ve Channels
// ==============================

// GÖREV 1: sayHello fonksiyonu
// 1 saniye bekle, sonra channel'a "Merhaba, [name]!" gönder
func sayHello(name string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Merhaba, %s!", name)
}

// GÖREV 2: sum fonksiyonu
// Slice'daki sayıların toplamını hesapla ve channel'a gönder
func sum(numbers []int, ch chan int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	ch <- total
}

// GÖREV 3: worker fonksiyonu
// jobs channel'dan iş al, işle, results channel'a sonuç gönder
// Her iş için: sonuç = iş * 2
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d işliyor: %d\n", id, job)
		results <- job * 2
	}
	// İpucu: for job := range jobs { ... }
	// Her işi aldığında "Worker X işliyor: Y" yazdır
}

// GÖREV 4: timeout örneği
// 2 saniye içinde channel'dan değer gelmezse "timeout" yazdır
func waitWithTimeout(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}

func main() {
	// Test 1: sayHello
	fmt.Println("=== Say Hello ===")
	helloCh := make(chan string)
	go sayHello("Go Developer", helloCh)
	msg := <-helloCh
	fmt.Println(msg)

	// Test 2: Parallel sum
	fmt.Println("\n=== Parallel Sum ===")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumCh := make(chan int)

	// İlk yarıyı bir goroutine'de topla
	go sum(numbers[:5], sumCh)
	// İkinci yarıyı başka goroutine'de topla
	go sum(numbers[5:], sumCh)

	// İki sonucu al ve topla
	sum1 := <-sumCh
	sum2 := <-sumCh
	fmt.Println("Toplam:", sum1+sum2) // 55 olmalı

	// Test 3: Worker pool
	fmt.Println("\n=== Worker Pool ===")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 3 worker başlat
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5 iş gönder
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Daha fazla iş yok

	// Sonuçları al
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Println("Sonuç:", result)
	}

	// Test 4: Timeout
	fmt.Println("\n=== Timeout ===")
	timeoutCh := make(chan string)
	// Hiçbir şey göndermiyoruz, timeout olmalı
	waitWithTimeout(timeoutCh)
}
