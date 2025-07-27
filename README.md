🔗 ShortLinker - сервис сокращения ссылок
Пет-проект на Go для преобразования длинных URL в компактные ссылки
Простой, но мощный инструмент с полным циклом обработки ссылок — от создания до аналитики.

🚀 Основные возможности
Молниеносное сокращение URL любой длины

Персонализированные короткие ссылки (по желанию пользователя)

Полноценная авторизация через JWT-токены

Детальная статистика по переходам

Гибкое API для интеграции с любым фронтендом

🛠 Технологический стек
go
import (
    "net/http"       // Нативный HTTP-сервер
    "github.com/go-playground/validator/v10" // Валидация
    "gorm.io/gorm"   // ORM для PostgreSQL
)
Ядро системы:

Go 1.20+ (чистый код без лишних зависимостей)

GORM — удобная работа с PostgreSQL

Validator — надежная валидация входящих данных

JWT — безопасная аутентификация

Docker — простота развертывания

⚡ Быстрый старт
Требования:
Docker 20.10+

Docker Compose 2.2+

bash
# 1. Клонируем репозиторий
git clone https://github.com/yourname/shortlinker.git
cd shortlinker

# 2. Запускаем сервисы
docker-compose up -d --build

# 3. Проверяем работу
curl http://localhost:8080/healthcheck
Сервис будет доступен на порту 8080

📊 Как это работает
Пользователь отправляет длинный URL

Сервис генерирует короткий код (например, xYz12)

Создается запись в PostgreSQL

При переходе по короткой ссылке:

Система находит оригинальный URL

Фиксирует факт перехода

Перенаправляет пользователя

🔐 Авторизация
Используем JWT-токены с сроком жизни 24 часа:

bash
# Получение токена
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'

# Ответ:
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
📝 Примеры API запросов
Создание короткой ссылки:

bash
curl -X POST http://localhost:8080/api/links \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/very/long/url"}'
Получение статистики:

bash
curl -X GET http://localhost:8080/api/links/stats \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
🗃 Структура базы данных
Основные таблицы PostgreSQL:

users — данные пользователей

links — соответствие коротких и длинных URL

clicks — статистика переходов
