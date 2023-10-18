# Simplebank

[Udemy "Backend Master Class [Golang + Postgres + Kubernetes + gRPC]"](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/) project course

## Progress:

- **Section 1: Working with Database [POSTGRES + SQLC]**
  - ✅ Design DB schema and generate SQL code with dbdiagram.io
  - ✅ Install & use Docker + Postgres + TablePlus to create DB 
  - ✅ How to write & run database migration in Golang
  - ✅ Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc
  - ✅ Write unit tests for database CRUD with random data in Golang
  - ✅ A clean way to implement database transaction in Golang
  - ✅ DB transaction lock & How to handle deadlock in Golang
  - ✅ How to avoid deadlock in DB transaction? Queries order matters!
  - ✅ Deeply understand transaction isolation levels & read phenomena
  - 🔨 Setup Github Actions for Golang + Postgres to run automated tests

- **Section 2: Building RESTful HTTP JSON API [Gin + JWT + PASETO]**
  - 🔲 Implement RESTful HTTP API in Go using Gin
  - 🔲 Load config from file & environment variables in Go with Viper
  - 🔲 Mock DB for testing HTTP API in Go and achieve 100% coverage
  - 🔲 Implement transfer money API with a custom params validator
  - 🔲 Add users table with unique & foreign key constraints in PostgreSQL
  - 🔲 How to handle DB errors in Golang correctly
  - 🔲 How to securely store passwords? Hash password in Go with Bcrypt!
  - 🔲 How to write stronger unit tests with a custom gomock matcher
  - 🔲 Why PASETO is better than JWT for token-based authentication?
  - 🔲 How to create and verify JWT & PASETO token in Golang
  - 🔲 Implement login user API that returns PASETO or JWT access token in Go
  - 🔲 Implement authentication middleware and authorization rules in Golang using Gin