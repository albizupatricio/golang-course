# Golang Course

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Exercises and projects from a Go course, in progressive order: from language fundamentals to a full REST API with authentication. Each numbered folder is an independent Go module (its own `go.mod`).

## Structure

| Folder | Content |
|---|---|
| 01 - Investiment Calculator | Fundamentals: variables, types, basic I/O |
| 02 - Basics Practice 1 | Fundamentals practice |
| 03 - Bank Practice 2 | Structs, file read/write (`fileops`) |
| 04 - Profit Calculator | Basic logic, file output |
| 05 - Pointers | Pointers |
| 06 - Structs | Structs, custom package (`user`) |
| 07 - Custom types | Custom types |
| 08 - Notes | Structs, JSON persistence (`note`, `todo`) |
| 09 - Lists-Arrays | Slices and arrays |
| 10 - Arrays Practice | Arrays practice |
| 11 - Functions | Functions |
| 12 - Recursion | Recursion |
| 13 - Variadic Functions | Variadic functions |
| 14 - Price Calculator | Concurrency with goroutines and channels (`cmdmanager`, `conversion`, `filemanager`, `iomanager`, `prices`) |
| 15 - Concurrency | Goroutines |
| 16 - Rest APIs | **Final project** — see below |

## Final project: 16 - Rest APIs

REST API built with [Gin](https://github.com/gin-gonic/gin), SQLite persistence, JWT authentication, and bcrypt-hashed passwords.

**Features:**
- `POST /signup`, `POST /login` — user signup and login
- `GET /events`, `GET /events/:id` — list and fetch events (public)
- `POST /events`, `PUT /events/:id`, `DELETE /events/:id` — create/update/delete events (requires authentication)
- `POST /events/:id/register`, `DELETE /events/:id/register` — register/cancel registration for an event (requires authentication)

**Running it:**

```bash
cd "16 - Rest APIs"
export JWT_SECRET="$(openssl rand -base64 32)"   # optional, falls back to a dev default if unset
go run main.go
```

The server listens on `localhost:8080`. The SQLite database (`api.db`) is created automatically on startup.

The `.http` files under `api-test/` (usable with VS Code's REST Client extension) contain example requests, including how to pass the JWT token on protected routes.

## Running any other exercise

```bash
cd "0X - Folder name"
go run .
```
