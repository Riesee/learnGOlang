package main

import (
	"fmt"
	"lessons/08_packages/mathutil"
)

func main() {
	// Test Add
	fmt.Println("5 + 3 =", mathutil.Add(5, 3)) // 8

	// Test Subtract
	fmt.Println("10 - 4 =", mathutil.Subtract(10, 4)) // 6

	// Test Square
	fmt.Println("7² =", mathutil.Square(7)) // 49

	// Test IsPrime
	fmt.Println("7 asal mı?", mathutil.IsPrime(7))   // true
	fmt.Println("10 asal mı?", mathutil.IsPrime(10)) // false
	fmt.Println("2 asal mı?", mathutil.IsPrime(2))   // true
	fmt.Println("1 asal mı?", mathutil.IsPrime(1))   // false

	// Bu satır HATA verir çünkü multiply private:
	// mathutil.multiply(2, 3)  // ❌ cannot refer to unexported name
}
