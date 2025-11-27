package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============================================
// EGZERSİZ 5: GORM ile Database İşlemleri
// ============================================
//
// GORM NEDİR?
// -----------
// Go için ORM (Object-Relational Mapping)
// SQL yazmadan database işlemi yaparsın
// TypeScript'teki Prisma, TypeORM gibi
//
// TEMEL KAVRAMLAR:
// ----------------
// Model     → Database tablosunu temsil eden struct
// Migration → Struct'tan tablo oluşturma
// Create    → INSERT
// Find      → SELECT
// Save      → UPDATE
// Delete    → DELETE
//
// BAĞLANTI:
// ---------
// dsn := "host=localhost user=postgres password=postgres dbname=taskdb port=5432"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//
// ============================================

// GÖREV 1: GORM ve PostgreSQL driver'ı import et
// "gorm.io/gorm"
// "gorm.io/driver/postgres"

// GÖREV 2: Task modeli tanımla
type Task struct {
	Title       string
	Description string
	Completed   bool
	gorm.Model
}

func main() {
	// ============================================
	// BÖLÜM 1: DATABASE BAĞLANTISI
	// ============================================
	dsn := "host=localhost user=postgres password=postgres dbname=taskdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database bağlantısı başarısız: " + err.Error())
	}
	// GÖREV 3: DSN (connection string) oluştur

	// GÖREV 4: Database'e bağlan

	fmt.Println("Database bağlantısı başarılı!")

	// ============================================
	// BÖLÜM 2: MIGRATION
	// ============================================

	db.AutoMigrate(&Task{})
	// GÖREV 5: Tabloyu oluştur (Auto Migration)

	fmt.Println("Migration tamamlandı!")

	task := Task{
		Title:       "Go öğren",
		Description: "GORM ile database",
		Completed:   false,
	}

	db.Create(&task)

	// ============================================
	// BÖLÜM 3: CRUD İŞLEMLERİ
	// ============================================

	// GÖREV 6: CREATE - Yeni task oluştur

	fmt.Println("Task oluşturuldu!")

	var tasks []Task
	db.Find(&tasks)

	// GÖREV 7: READ - Tüm taskları getir

	fmt.Println("Tüm tasklar:")

	for _, task := range tasks {
		fmt.Printf("- %d: %s\n", task.ID, task.Title)
	}

	var singleTask Task

	db.First(&singleTask, 8)

	db.Model(&singleTask).Updates(Task{Title: "Yeni başlık", Completed: true})
	// GÖREV 8: READ - Tek task getir (ID ile)

	// GÖREV 9: UPDATE - Task güncelle

	fmt.Println("Task güncellendi!")

	// GÖREV 10: DELETE - Task sil (soft delete)
	var deleteTask Task
	db.First(&deleteTask, 7)

	db.Delete(&deleteTask)

	fmt.Println("Task silindi!")

	// ============================================
	// BONUS: WHERE ile sorgulama
	// ============================================
	var completedTasks []Task
	db.Where("completed = ?", true).Find(&completedTasks)
	fmt.Printf("Tamamlanan task sayısı: %d\n", len(completedTasks))

	var searchTasks []Task
	db.Where("title LIKE ?", "%Go%").Find(&searchTasks)
	fmt.Println("'Go' içeren tasklar:")
	for _, t := range searchTasks {
		fmt.Printf("- %s\n", t.Title)
	}
}

// ============================================
// ÇALIŞTIRMAK İÇİN:
// ============================================
// 1. docker-compose up -d (PostgreSQL başlat)
// 2. go run main.go
// ============================================
