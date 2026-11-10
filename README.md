# who-has-my-book

> Keep track of the books you lent to your friends.

A personal app to manage a home library: register books (with covers), friends, loans with due dates and returns — keeping the full loan history.

- **Admin area** (`/admin`): manage books, friends, loans and returns. Google sign-in, restricted to an allowlisted email.
- **Friend area** (`/friend/<token>`): each friend gets a unique link to see which books are with them.

## Stack

| Layer    | Tech                                                    |
| -------- | ------------------------------------------------------- |
| Backend  | Go (`net/http`), GORM, goose migrations, `log/slog`     |
| Database | PostgreSQL                                              |
| Frontend | Next.js (App Router), React, TypeScript, Tailwind CSS   |
| Storage  | Cloudflare R2 (book covers), Open Library Covers API as fallback |

## Project structure

```
.
├── backend/            # Go API
│   ├── cmd/api/        # entrypoint
│   ├── internal/       # domains (book, friend, loan) + platform (config, db, server...)
│   └── migrations/     # goose SQL migrations
├── frontend/           # Next.js app
└── docker-compose.yml  # local PostgreSQL
```

## Getting started

Requirements: Go 1.24+, Node 20+, Docker.

```bash
# 1. Env vars
cp .env.example backend/.env

# 2. Database
docker compose up -d

# 3. Backend (http://localhost:8080)
cd backend
go run ./cmd/api

# 4. Frontend (http://localhost:3000)
cd frontend
npm install
npm run dev
```

Health check: `GET http://localhost:8080/health`

## License

[MIT](LICENSE)
