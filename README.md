# 🧑‍🎓 Student Management API

A simple RESTful API built in Go for managing student records using SQLite. The API supports creating, retrieving, listing, and deleting student data. It features request validation, structured logging, and graceful server shutdown.

---
- Command to run project  -  go run cmd/students-api/main.go -config config/local.yaml
## 🚀 Features

- Create new student records
- Get a student by ID
- Get all students
- Delete a student by ID
- Input validation using `go-playground/validator/v10`
- Clean and modular code structure
- Graceful shutdown on termination signals

---

## 🛠 Tech Stack

- **Language:** Go
- **Database:** SQLite
- **Validation:** go-playground/validator
- **Routing:** `http.ServeMux` (Go standard library)


---



