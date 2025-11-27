package main

import "fmt"

// DERS 4: Structs ve Methods
// ==========================

// GÖREV 1: Rectangle struct'ı tanımla
// Width ve Height alanları olsun (float64)
type Rectangle struct {
	Width, Height float64
}

// GÖREV 2: Area methodu yaz
// Dikdörtgenin alanını hesapla (Width * Height)
// Value receiver kullan
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// GÖREV 3: Scale methodu yaz
// Dikdörtgenin boyutlarını verilen faktörle çarp
// Pointer receiver kullanmalısın (orijinali değiştirmeli)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// GÖREV 4: BankAccount struct'ı ve methodları
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit methodu - hesaba para yatır
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

// Withdraw methodu - hesaptan para çek
// Yetersiz bakiye varsa false döndür
func (b *BankAccount) Withdraw(amount float64) bool {
	if b.Balance >= amount {
		b.Balance -= amount
		return true
	}
	return false
}

func main() {
	// Test Rectangle
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Alan:", rect.Area()) // 50 olmalı

	rect.Scale(2)
	fmt.Println("Scale sonrası alan:", rect.Area()) // 200 olmalı

	// Test BankAccount
	account := BankAccount{Owner: "Ali", Balance: 1000}
	account.Deposit(500)
	fmt.Println("Bakiye:", account.Balance) // 1500 olmalı

	success := account.Withdraw(2000)
	fmt.Println("2000 çekme başarılı mı?", success) // false olmalı
	fmt.Println("Bakiye:", account.Balance)         // 1500 olmalı (değişmemeli)

	success = account.Withdraw(500)
	fmt.Println("500 çekme başarılı mı?", success) // true olmalı
	fmt.Println("Bakiye:", account.Balance)        // 1000 olmalı
}
