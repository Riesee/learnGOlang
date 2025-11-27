package main

import (
	"errors"
	"fmt"
)

// DERS 6: Error Handling
// ======================

// Önceden tanımlı hatalar (sentinel errors)
var (
	ErrDivideByZero = errors.New("sıfıra bölme hatası")
	ErrNegativeAge  = errors.New("yaş negatif olamaz")
	ErrEmptyName    = errors.New("isim boş olamaz")
)

// GÖREV 1: safeDivide fonksiyonu
// Sıfıra bölme durumunda ErrDivideByZero döndür
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// GÖREV 2: Person struct ve validate fonksiyonu
type Person struct {
	Name string
	Age  int
}

// Validate metodu - Person'ı doğrula
// - Name boşsa ErrEmptyName döndür
// - Age negatifse ErrNegativeAge döndür
// - Her şey doğruysa nil döndür
func (p Person) Validate() error {
	switch {
	case p.Name == "":
		return ErrEmptyName
	case p.Age < 0:
		return ErrNegativeAge
	}
	return nil
}

// GÖREV 3: Custom error tipi
// InsufficientFundsError - yetersiz bakiye hatası
type InsufficientFundsError struct {
	Requested float64
	Available float64
}

// Error metodu - error interface'ini implement et
func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("Yetersiz bakiye: %.2f istendi, %.2f mevcut", e.Requested, e.Available)
}

// GÖREV 4: withdraw fonksiyonu
// Yetersiz bakiyede InsufficientFundsError döndür
func withdraw(balance, amount float64) (float64, error) {
	if balance < amount {
		return balance, InsufficientFundsError{Requested: amount, Available: balance}
	}
	return balance - amount, nil
}

func main() {
	// Test 1: safeDivide
	fmt.Println("=== Safe Divide ===")
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Hata:", err)
		// Sentinel error kontrolü
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("(Bu bir sıfıra bölme hatasıydı)")
		}
	}

	// Test 2: Person Validate
	fmt.Println("\n=== Person Validate ===")
	p1 := Person{Name: "Ali", Age: 25}
	if err := p1.Validate(); err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("p1 geçerli")
	}

	p2 := Person{Name: "", Age: 25}
	if err := p2.Validate(); err != nil {
		fmt.Println("Hata:", err)
	}

	p3 := Person{Name: "Ayşe", Age: -5}
	if err := p3.Validate(); err != nil {
		fmt.Println("Hata:", err)
	}

	// Test 3: withdraw
	fmt.Println("\n=== Withdraw ===")
	newBalance, err := withdraw(1000, 500)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Yeni bakiye:", newBalance)
	}

	newBalance, err = withdraw(100, 500)
	if err != nil {
		fmt.Println("Hata:", err)
		// Custom error tipini kontrol et
		var insufficientErr InsufficientFundsError
		if errors.As(err, &insufficientErr) {
			fmt.Printf("(İstenen: %.2f, Mevcut: %.2f)\n",
				insufficientErr.Requested, insufficientErr.Available)
		}
	}
}
