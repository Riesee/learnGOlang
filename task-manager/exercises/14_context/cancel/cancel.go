package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d iptal edildi\n", id)
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(500 *time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("\n--- İptal ediliyor ---")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\n--- İptal edildi ---")
}