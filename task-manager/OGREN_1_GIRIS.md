# 🎯 Ders 1: Go Projesi Nasıl Çalışır?

## go.mod Nedir?

`go.mod` = Node.js'teki `package.json`

```go
module task-manager     // Projenin adı (import'larda kullanılır)
go 1.21                 // Go versiyonu
require (...)           // Bağımlılıklar (dependencies)
```

## main.go Nedir?

Her Go programı `main` package'ından ve `main()` fonksiyonundan başlar.

```go
package main  // Bu dosya çalıştırılabilir bir program

func main() {
    // Program buradan başlar
}
```

## internal/ Klasörü Nedir?

Go'da `internal/` özel bir klasör. Bu klasördeki kodlar SADECE bu proje içinden import edilebilir.
Dışarıdan kimse `import "task-manager/internal/..."` yapamaz.

Bu güvenlik için - private kod gibi düşün.

---

## 🏋️ EGZERSİZ 1

`exercises/01_hello/main.go` dosyasını aç ve tamamla.
