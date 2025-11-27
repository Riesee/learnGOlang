package main

import "fmt"

// DERS 2: Döngüler ve Kontrol Yapıları
// ====================================

// GÖREV 1: FizzBuzz klasiği
// 1'den n'e kadar sayıları yazdır:
// - 3'e bölünüyorsa "Fizz"
// - 5'e bölünüyorsa "Buzz"
// - Her ikisine de bölünüyorsa "FizzBuzz"
// - Hiçbiri değilse sayının kendisi
func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// GÖREV 2: Slice'daki sayıların toplamını bul
// İpucu: for range kullan
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// GÖREV 3: Bir sayının faktöriyelini hesapla
// Örnek: factorial(5) = 5 * 4 * 3 * 2 * 1 = 120
func factorial(n int) int {
	var result = 1
	for i := 2 ; i <= n ; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println("=== FizzBuzz (1-15) ===")
	fizzBuzz(15)

	fmt.Println("\n=== Sum ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Toplam:", sum(nums)) // 15 olmalı

	fmt.Println("\n=== Factorial ===")
	fmt.Println("5! =", factorial(5))   // 120 olmalı
	fmt.Println("0! =", factorial(0))   // 1 olmalı
}
