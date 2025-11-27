package main

import "fmt"

// DERS 1: Temel Syntax
// ====================

// GÖREV 1: Bu fonksiyonu tamamla
// İki sayıyı çarp ve sonucu döndür
func multiply(a int, b int) int {
	return a * b 
}

// GÖREV 2: Bu fonksiyonu tamamla
// Bir sayının karesini al, eğer negatifse hata döndür
// İpucu: divide fonksiyonundaki pattern'i kullan
func square(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("negatif sayı girdiniz")
	}
	return n * n, nil
}

func main() {
	// Test 1: multiply
	result := multiply(4, 5)
	fmt.Println("4 x 5 =", result) // 20 olmalı

	// Test 2: square
	sq, err := square(5)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("5'in karesi =", sq) // 25 olmalı
	}

	// Test 3: square negatif sayı ile
	sq2, err2 := square(-3)
	if err2 != nil {
		fmt.Println("Hata:", err2) // Hata mesajı görmeli
	} else {
		fmt.Println("-3'ün karesi =", sq2)
	}
}
