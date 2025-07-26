# ShortURL

**Short URL** — это сервис для сокращения ссылок, написанный на Go. Он позволяет создавать короткие URL-адреса, которые перенаправляют на оригинальные длинные ссылки.

## 🛠️ Установка и запуск

1. **Клонируйте репозиторий:**
```bash
git clone https://github.com/Mutter0815/short-url.git
cd short-url

```
2. **Настройте переменные окружения:**

Создайте файл `.env` в корне проекта и добавьте необходимые переменные. Пример содержимого:

```env
DB_USER="postgres"
DB_HOST="localhost"
DB_PORT="5433"
DB_NAME="person_db"
DB_PASSWORD="admin"
```
3. **Запусти через Docker**
```bash
docker-compose up --build
```

Сервис доступен на [localhost:8080](http://localhost:8080)


**Стек технологий**

Go 1.23

PostgreSQL 15

Gin (HTTP сервер)

Docker + Docker Compose


## 📌 Пример использования API
Создание короткой ссылки:
Запрос:

```http
POST /link
Content-Type: application/json

{
  "link": "https://example.com/very/long/url"
}

```

Ответ: 
```json
{
  "link": "http://localhost:8080/abc123"
}

```

Переход по короткой ссылке:
Запрос:

```html
GET /abc123
```
Происходит перенаправление на оригинальный URL: 
`https://example.com/very/long/url`.

Получение оригинального URL по короткой ссылке
Запрос:

```html
GET /orginalurl?shortlink=abc123
```
Ответ:
```json
{
  "originalURL": "https://example.com/very/long/url"
}
```

Получение короткой ссылки по оригинальному URL:

Запрос:
```http
GET /shortlink?originalurl=https://example.com/very/long/url
```
**Ответ:**
```json
{
  "shortLink": "http://localhost:8080/abc123"
}
```
