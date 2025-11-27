# Task Manager API

Go ile yazılmış production-ready REST API.

## Öğrenilecekler
- [x] Proje yapısı
- [ ] Gin Gonic (REST API)
- [ ] Viper (Configuration)
- [ ] Zap (Logging)
- [ ] GORM + PostgreSQL
- [ ] JWT Authentication
- [ ] Docker

## Çalıştırma

```bash
# Development
go run cmd/api/main.go

# Docker ile
docker-compose up
```

## API Endpoints

```
POST   /api/auth/register  - Kayıt ol
POST   /api/auth/login     - Giriş yap
GET    /api/tasks          - Tüm tasklar (auth gerekli)
POST   /api/tasks          - Task oluştur (auth gerekli)
GET    /api/tasks/:id      - Tek task
PUT    /api/tasks/:id      - Task güncelle
DELETE /api/tasks/:id      - Task sil
```
