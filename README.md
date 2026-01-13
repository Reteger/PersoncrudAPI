# Person CRUD API 🚀

Golang REST API для полного CRUD-цикла работы с сущностью "Person" (email, phone, имя, фамилия). PostgreSQL + Docker. Готово к запуску за 1 команду!

[![GitHub Repo stars](https://img.shields.io/github/stars/Reteger/PersoncrudAPI?style=social)](https://github.com/Reteger/PersoncrudAPI)
[![GitHub issues](https://img.shields.io/github/issues/Reteger/PersoncrudAPI)](https://github.com/Reteger/PersoncrudAPI/issues)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-brightgreen.svg)](https://golang.org)

##  Содержание
- [Быстрый запуск](#быстрый-запуск)
- [Тестирование API](#тестирование-api)
- [CRUD операции](#crud-операции)
- [Автотесты](#автотесты)
- [Управление БД](#управление-бд)

## Быстрый запуск

###  Способ 1: Docker (рекомендуется)
```bash
git clone https://github.com/Reteger/PersoncrudAPI.git
cd PersoncrudAPI
docker-compose up --build
```

# БД
``` docker-compose up -d postgres ```

# App (нужен Go 1.21+)

```
go run cmd/app/main.go
Тестирование API
 Health Check (PowerShell)
powershell
Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
```
Ответ: {"status":"ok"}

CRUD операции
# 1. Получить всех
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET
```
Пример:

json
[
  {
    "id": 1,
    "email": "ivan@example.com",
    "phone": "+79161234567",
    "firstName": "Ivan",
    "lastName": "Petrov"
  }
]

# 2. Создать
```powershell
$body = '{"email": "newuser@example.com", "phone": "+79161234567", "firstName": "Alex", "lastName": "Smith"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method POST -ContentType "application/json" -Body $body
```
Ответ: {"id": 2, "message": "Person created successfully"}

# 3. По ID
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method GET
```
# 4. Обновить
```powershell
$body = '{"email": "ivan.updated@example.com", "phone": "+79169876543", "firstName": "Ivan", "lastName": "Ivanov"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method PUT -ContentType "application/json" -Body $body
```
Ответ: {"message": "Person updated successfully"}

# 5. Удалить
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method DELETE
```
Ответ: {"message": "Person deleted successfully"}

# Автотесты
Полный CRUD-тест:

```powershell
Write-Host "=== ПОЛНЫЙ ТЕСТ CRUD ===" -ForegroundColor Green
```
# 1. Health
```
$health = Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
Write-Host " Статус: $($health.status)" -ForegroundColor Green
```
# 2. Create
```
$body = '{"email": "test@example.com", "phone": "+79161234567", "firstName": "Test", "lastName": "User"}'
$result = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method POST -ContentType "application/json" -Body $body
$personId = $result.id
Write-Host "✅ Создан ID: $personId" -ForegroundColor Green
```
# 3. Read all
```
$persons = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET
Write-Host "✅ Записей: $($persons.Count)" -ForegroundColor Green
```
# 4. Read by ID
```
$person = Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method GET
Write-Host "✅ Email: $($person.email)" -ForegroundColor Green
```
# 5. Update
```
$updateBody = '{"email": "updated@example.com", "phone": "+79169876543", "firstName": "Updated", "lastName": "User"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method PUT -ContentType "application/json" -Body $updateBody
Write-Host "✅ Обновлено" -ForegroundColor Green
```
# 6. Delete
```
Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method DELETE
Write-Host "✅ Удалено" -ForegroundColor Green

Write-Host "🎉 ТЕСТ ПРОВЕРЕН!" -ForegroundColor Green
Запуск: powershell -ExecutionPolicy Bypass -File test-crud.ps1
```
Управление БД
```bash
# Остановить + очистить
docker-compose down -v
```

# Перезапуск БД
docker-compose up -d postgres
