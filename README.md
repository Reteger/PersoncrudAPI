# 📋 Инструкция по запуску и использованию Person CRUD API
# 🚀 Быстрый запуск
# Способ 1: Docker (рекомендуется)

Клонировать проект

git clone https://github.com/Reteger/PersoncrudAPI.git
cd personcrud

Запустить всё одной командой

docker-compose up --build

# Способ 2: Локальный запуск

Запустить базу данных

docker-compose up -d postgres

Запустить приложение

go run cmd/app/main.go

Приложение будет доступно по адресу: http://localhost:8080
________________________________________
# 🧪 Тестирование API
Для PowerShell:

1. Проверка здоровья сервиса

Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET

Результат:{"status":"ok"}

________________________________________
# 👥 РАБОТА С ЛЮДЬМИ
📋 1. ПОЛУЧЕНИЕ ВСЕХ ЛЮДЕЙ

 Получить список всех людей
 
Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET

Пример ответа:

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
________________________________________
# ➕ 2. ДОБАВЛЕНИЕ НОВОГО ЧЕЛОВЕКА

 Добавить нового человека
 
$body = '{"email": "newuser@example.com", "phone": "+79161234567", "firstName": "Alex", "lastName": "Smith"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method POST -ContentType "application/json" -Body $body

Пример ответа:

json
{
  "id": 2,
  "message": "Person created successfully"
}
________________________________________
# 👤 3. ПОЛУЧЕНИЕ ЧЕЛОВЕКА ПО ID

 Получить информацию о человеке с ID=1
 
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method GET

Пример ответа:

json
{
  "id": 1,
  "email": "ivan@example.com",
  "phone": "+79161234567",
  "firstName": "Ivan",
  "lastName": "Petrov"
}
________________________________________
# ✏️ 4. ОБНОВЛЕНИЕ ДАННЫХ ЧЕЛОВЕКА

Обновить данные человека с ID=1
 
$body = '{"email": "ivan.updated@example.com", "phone": "+79169876543", "firstName": "Ivan", "lastName": "Ivanov"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method PUT -ContentType "application/json" -Body $body

Пример ответа:

json
{
  "message": "Person updated successfully"
}
________________________________________
# 🗑️ 5. УДАЛЕНИЕ ЧЕЛОВЕКА

 Удалить человека с ID=1

Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method DELETE

Пример ответа:

json
{
  "message": "Person deleted successfully"
}

________________________________________
# 🔍 ПРОВЕРКА РЕЗУЛЬТАТОВ

Просмотр всех записей после операций:

 Посмотреть текущее состояние базы

$persons = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET Write-Host "Всего записей: $($persons.Count)" $persons | Format-Table id, email, firstName, lastName -AutoSize

# 🎯 АВТОМАТИЧЕСКОЕ ТЕСТИРОВАНИЕ

Полный тест всех операций:

powershell

Write-Host "=== ПОЛНЫЙ ТЕСТ CRUD ОПЕРАЦИЙ ===" -ForegroundColor Green

#1. Проверка здоровья
   
Write-Host "1. Проверка здоровья..." -ForegroundColor Yellow
$health = Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
Write-Host "Статус: $($health.status)" -ForegroundColor Green

#2. Создание

Write-Host "2. Создание человека..." -ForegroundColor Yellow
$body = '{"email": "test@example.com", "phone": "+79161234567", "firstName": "Test", "lastName": "User"}'
$result = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method POST -ContentType "application/json" -Body $body
$personId = $result.id
Write-Host "Создан с ID: $personId" -ForegroundColor Green

#3. Чтение всех
   
Write-Host "3. Чтение всех записей..." -ForegroundColor Yellow
$persons = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET
Write-Host "Найдено записей: $($persons.Count)" -ForegroundColor Green

#4. Чтение по ID
   
Write-Host "4. Чтение по ID..." -ForegroundColor Yellow
$person = Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method GET
Write-Host "Email: $($person.email)" -ForegroundColor Green

#5. Обновление
   
Write-Host "5. Обновление..." -ForegroundColor Yellow
$updateBody = '{"email": "updated@example.com", "phone": "+79169876543", "firstName": "Updated", "lastName": "User"}'
Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method PUT -ContentType "application/json" -Body $updateBody
Write-Host "Данные обновлены" -ForegroundColor Green

#6. Удаление

Write-Host "6. Удаление..." -ForegroundColor Yellow
Invoke-RestMethod -Uri "http://localhost:8080/persons/$personId" -Method DELETE
Write-Host "Запись удалена" -ForegroundColor Green

#7. Проверка удаления

Write-Host "7. Проверка удаления..." -ForegroundColor Yellow
$persons = Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method GET
Write-Host "Осталось записей: $($persons.Count)" -ForegroundColor Green

Write-Host "=== ТЕСТ ЗАВЕРШЕН ===" -ForegroundColor Green

# 🗂️ УПРАВЛЕНИЕ БАЗОЙ ДАННЫХ
Очистка всей базы данных:
powershell

 #Перейти в папку проекта
 
cd C:\Users\esnas\GolandProjects\personcrud

 #Остановить и очистить базу
 
docker-compose down -v
docker-compose up -d postgres

