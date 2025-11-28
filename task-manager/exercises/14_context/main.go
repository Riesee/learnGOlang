package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("✅ İş tamamlandı!")
        return nil
	case <-ctx.Done():
		fmt.Println("❌ İş iptal edildi:", ctx.Err())
        return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("İş başlıyor...")
    err := slowOperation(ctx)
    if err != nil {
        fmt.Println("Hata:", err)
    }

	fmt.Println("İş bitti")
}