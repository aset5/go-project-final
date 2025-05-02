
Spring boot + Go + PostgreSQL жобасы.  
Бұл жоба RESTful API құрайды, GORM, Gin және миграциялармен жұмыс істейді.


##  Технологиялар

[Go](https://golang.org/)
[Gin](https://github.com/gin-gonic/gin)
[GORM](https://gorm.io/)
[PostgreSQL](https://www.postgresql.org/)
[golang-migrate](https://github.com/golang-migrate/migrate)


## Қолдану нұсқаулығы

### 1. Репозиторийді клондау

https://github.com/aset5/springProject-Go-code.git
### 2. Миграция жасау

```bash
migrate -path migrations -database "postgres://postgres:пароль@localhost:5432/postgres?sslmode=disable" up
```

### 3. Серверді іске қосу

```bash
go run main.go
```


## API маршруты (Мысалы)

| HTTP әдісі | Маршрут         | Сипаттама        |
|------------|------------------|------------------|
| GET        | `/products`      | Барлық өнімдер   |
| POST       | `/products`      | Жаңа өнім қосу   |
| GET        | `/products/:id`  | Өнімді көру      |
| PUT        | `/products/:id`  | Өнімді жаңарту   |
| DELETE     | `/products/:id`  | Өнімді жою       |

---

## Жоба құрылымы

```bash
.
├── cmd/                # main.go
├── config/             # DB конфиг
├── models/             # Модельдер (Product)
├── controllers/        # Логика
├── routes/             # API маршруттары
├── migrations/         # SQL миграциялар
├── .env                # Құпия деректер
├── docker-compose.yml
├── go.mod
├── README.md
```

---


