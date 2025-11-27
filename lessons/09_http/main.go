package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// DERS 9: HTTP ve REST API
// ========================

// Todo struct - JSON tag'lerine dikkat!
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// In-memory database (gerçek projede DB kullanılır)
var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex // concurrent access için
)

// GÖREV 1: getAllTodos handler
// GET /todos - tüm todo'ları JSON olarak döndür
func getAllTodos(w http.ResponseWriter, r *http.Request) {
	// BURAYA KOD YAZ
	// 1. Content-Type header'ı set et: "application/json"
	// 2. todos slice'ını JSON olarak encode et ve yaz
}

// GÖREV 2: createTodo handler
// POST /todos - yeni todo oluştur
func createTodo(w http.ResponseWriter, r *http.Request) {
	// BURAYA KOD YAZ
	// 1. Request body'den Todo'yu decode et
	// 2. ID ata (nextID kullan, sonra artır)
	// 3. todos slice'a ekle
	// 4. Status 201 (Created) döndür
	// 5. Oluşturulan todo'yu JSON olarak döndür
	// İpucu: mu.Lock() ve mu.Unlock() kullan
}

// GÖREV 3: getTodoByID handler
// GET /todos/{id} - belirli bir todo'yu döndür
func getTodoByID(w http.ResponseWriter, r *http.Request) {
	// BURAYA KOD YAZ
	// 1. URL'den id'yi al: r.PathValue("id")
	// 2. String'i int'e çevir: strconv.Atoi()
	// 3. todos içinde bul
	// 4. Bulunamazsa 404 döndür
	// 5. Bulunursa JSON olarak döndür
}

// GÖREV 4: deleteTodo handler
// DELETE /todos/{id} - todo'yu sil
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// BURAYA KOD YAZ
	// 1. URL'den id'yi al
	// 2. todos içinde bul ve sil
	// 3. Bulunamazsa 404 döndür
	// 4. Başarılıysa 204 (No Content) döndür
}

// Helper: JSON response gönder
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Helper: Error response gönder
func errorResponse(w http.ResponseWriter, status int, message string) {
	jsonResponse(w, status, map[string]string{"error": message})
}

func main() {
	// Route tanımları (Go 1.22+ syntax)
	http.HandleFunc("GET /todos", getAllTodos)
	http.HandleFunc("POST /todos", createTodo)
	http.HandleFunc("GET /todos/{id}", getTodoByID)
	http.HandleFunc("DELETE /todos/{id}", deleteTodo)

	fmt.Println("Server başlatılıyor: http://localhost:8080")
	fmt.Println("\nTest komutları:")
	fmt.Println("  curl http://localhost:8080/todos")
	fmt.Println("  curl -X POST -d '{\"title\":\"Go öğren\"}' http://localhost:8080/todos")
	fmt.Println("  curl http://localhost:8080/todos/1")
	fmt.Println("  curl -X DELETE http://localhost:8080/todos/1")

	http.ListenAndServe(":8080", nil)
}
