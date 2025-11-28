# Exercise 11: GraphQL API

Bu egzersizde GraphQL API geliştirmeyi öğreniyoruz. 07_fullstack'teki PostgreSQL veritabanını kullanıyoruz.

## 📁 Proje Yapısı

```
11_graphql/
├── main.go                      # Uygulama giriş noktası
├── config.yaml                  # DB konfigürasyonu
├── gqlgen.yml                   # gqlgen kod üretici ayarları
├── graph/
│   ├── schema.graphqls          # Ana GraphQL şeması
│   ├── task.graphqls            # Task şeması
│   ├── user.graphqls            # User şeması
│   ├── generated/
│   │   └── generated.go         # ⚠️ Otomatik üretilir - DOKUNMA
│   ├── model/
│   │   └── models_gen.go        # ⚠️ Otomatik üretilir - DOKUNMA
│   └── resolverfn/
│       ├── resolver.go          # Dependency injection (DB)
│       ├── schema.resolvers.go  # Resolver bağlantıları
│       ├── task.resolvers.go    # ✅ Task iş mantığı
│       └── user.resolvers.go    # ✅ User iş mantığı
└── internal/
    ├── config/                  # Config yükleme
    ├── database/                # DB bağlantısı
    └── model/                   # GORM modelleri (DB)
```

## 🚀 Çalıştırma

1. PostgreSQL'in çalıştığından emin ol (07_fullstack'ten):
```bash
cd ../07_fullstack
docker-compose up -d
```

2. GraphQL sunucusunu başlat:
```bash
cd ../11_graphql
go run main.go
```

3. Tarayıcıda aç: http://localhost:8080

## 📝 Örnek Sorgular

### Tüm Task'ları Listele
```graphql
query {
  tasks {
    id
    title
    description
    completed
    userId
    createdAt
  }
}
```

### Tek Task Getir
```graphql
query {
  task(id: "1") {
    id
    title
    completed
  }
}
```

### Yeni Task Oluştur
```graphql
mutation {
  createTask(input: {
    title: "GraphQL öğren"
    description: "gqlgen ile API yaz"
  }) {
    id
    title
    createdAt
  }
}
```

### Task Güncelle
```graphql
mutation {
  updateTask(id: "1", input: {
    completed: true
  }) {
    id
    title
    completed
  }
}
```

### Task Sil
```graphql
mutation {
  deleteTask(id: "1")
}
```

### Kullanıcıları ve Task'larını Getir
```graphql
query {
  users {
    id
    name
    email
    tasks {
      id
      title
      completed
    }
  }
}
```

## 🔑 Önemli Kavramlar

### 1. Schema (.graphqls)
API'nin sözleşmesi. Tipler, Query ve Mutation'lar burada tanımlanır.

### 2. Resolver
Schema'daki her field için çalışan fonksiyon. İş mantığı burada.

### 3. gqlgen
Schema'dan Go kodu üreten araç. `go run github.com/99designs/gqlgen generate`

### 4. Query vs Mutation
- **Query**: Okuma işlemleri (GET)
- **Mutation**: Yazma işlemleri (POST/PUT/DELETE)

## 🔄 Geliştirme Akışı

1. `.graphqls` dosyasında şema değişikliği yap
2. `go run github.com/99designs/gqlgen generate` çalıştır
3. Yeni resolver'ları doldur (panic → gerçek kod)
4. Test et
