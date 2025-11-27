package mathutil

// DERS 8: Packages
// ================

// GÖREV 1: Add fonksiyonu (Public)
// İki sayıyı topla
func Add(a, b int) int {
	return a + b
}

// GÖREV 2: Subtract fonksiyonu (Public)
// İlk sayıdan ikinciyi çıkar
func Subtract(a, b int) int {
	return a - b
}

// GÖREV 3: multiply fonksiyonu (private - küçük harf!)
// İki sayıyı çarp - sadece bu package içinden erişilebilir
func multiply(a, b int) int {
	return a * b
}

// GÖREV 4: Square fonksiyonu (Public)
// Bir sayının karesini al
// İpucu: multiply fonksiyonunu kullan
func Square(n int) int {
	return multiply(n, n)
}

// GÖREV 5: IsPrime fonksiyonu (Public)
// Bir sayının asal olup olmadığını kontrol et
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
