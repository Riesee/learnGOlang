package service

import (
	"fmt"
	"sync"
	"time"
)

func SendNotification(userID uint, message string) error {
    fmt.Printf("📧 User %d'e gönderiliyor: %s\n", userID, message)
	time.Sleep(1 * time.Second)
	fmt.Printf("✅ User %d'e gönderildi\n", userID)
	return nil	
}

func SendNotificationToMany(userIDs []uint, message string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(userIDs))
	for _, userID := range userIDs {
		go func(id uint) {
			defer wg.Done()
			SendNotification(id, message)
		}(userID)
	}
	wg.Wait()
	return nil
}