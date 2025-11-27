package main

import "fmt"

// DERS 3: Pointers
// ================

// GÖREV 1: swap fonksiyonu
// İki değişkenin değerlerini yer değiştir
// İpucu: Pointer kullanmalısın, yoksa çalışmaz
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// GÖREV 2: increment fonksiyonu
// Verilen sayıyı 1 artır (orijinal değeri değiştirmeli)
func increment(n *int) {
	*n = *n + 1
}

// GÖREV 3: setToZero fonksiyonu
// Slice'daki tüm elemanları 0 yap
// NOT: Slice zaten reference gibi davranır, pointer gerekmez
func setToZero(nums []int) {
	for i := range nums {
		nums[i] = 0
	}
}

func main() {
	// Test 1: swap
	x, y := 5, 10
	fmt.Printf("Önce: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Sonra: x=%d, y=%d\n", x, y) // x=10, y=5 olmalı

	// Test 2: increment
	num := 42
	increment(&num)
	fmt.Println("Increment:", num) // 43 olmalı

	// Test 3: setToZero
	numbers := []int{1, 2, 3, 4, 5}
	setToZero(numbers)
	fmt.Println("SetToZero:", numbers) // [0 0 0 0 0] olmalı
}
