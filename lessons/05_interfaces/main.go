package main

import (
	"fmt"
	"math"
)

// DERS 5: Interfaces
// ==================

// Shape interface'i - Area ve Perimeter metodları olmalı
type Shape interface {
	Area() float64
	Perimeter() float64
}

// GÖREV 1: Rectangle struct'ı (önceki dersten)
type Rectangle struct {
	Width, Height float64
}

// Rectangle için Area metodu
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle için Perimeter (çevre) metodu
func (r Rectangle) Perimeter() float64 {
	return ( r.Width + r.Height ) * 2
}

// GÖREV 2: Circle struct'ı
type Circle struct {
	Radius float64
}

// Circle için Area metodu (π * r²)
// İpucu: math.Pi kullan
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle için Perimeter metodu (2 * π * r)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// GÖREV 3: Triangle struct'ı
type Triangle struct {
	A, B, C float64 // üç kenar uzunluğu
}

// Triangle için Area metodu
// Heron formülü: √(s(s-a)(s-b)(s-c)) where s = (a+b+c)/2
// İpucu: math.Sqrt kullan
func (t Triangle) Area() float64 {
	area := (t.A + t.B + t.C) / 2
	return math.Sqrt(area * (area - t.A) * (area - t.B) * (area - t.C))
}

// Triangle için Perimeter metodu
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// GÖREV 4: PrintShapeInfo fonksiyonu
// Shape interface'i alan ve bilgilerini yazdıran fonksiyon
func PrintShapeInfo(s Shape) {
	fmt.Println("Alan:", s.Area(), "Çevre:", s.Perimeter())
}

// GÖREV 5: TotalArea fonksiyonu
// Birden fazla Shape alıp toplam alanı döndür
// İpucu: ...Shape (variadic parameter) kullan
func TotalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}
	triangle := Triangle{A: 3, B: 4, C: 5}

	fmt.Println("=== Shape Bilgileri ===")
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
	PrintShapeInfo(triangle)

	fmt.Println("\n=== Toplam Alan ===")
	total := TotalArea(rect, circle, triangle)
	fmt.Printf("Toplam alan: %.2f\n", total)
	// Beklenen: 50 + 28.27 + 6 = ~84.27
}
