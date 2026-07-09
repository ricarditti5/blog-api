# Blog API

A lightweight REST API for a blogging platform, built with Go's standard library and PostgreSQL. Supports full CRUD operations on posts with dynamic search and filtering.

Based on the [Roadmap.sh Blogging Platform API](https://roadmap.sh/projects/blogging-platform-api) project.

## Tech Stack

| Layer | Technology |
|---|---|
| **Language** | Go 1.26.4 |
| **HTTP** | `net/http` (Go 1.22+ pattern routing) |
| **Database** | PostgreSQL via `pgx/v5` (connection pooling with `pgxpool`) |
| **Configuration** | `spf13/viper` + `subosito/gotenv` |

## Project Structure

```
blog-api/
├── cmd/api/main.go               # Entry point
├── internal/
│   ├── config/                    # Config loading (.env + viper) + DB connection
│   ├── handlers/                  # HTTP request handlers
│   ├── models/                    # Data structures (Posts, PostFilter)
│   ├── repository/                # Database queries (CRUD + search)
│   ├── routes/                    # Route registration
│   ├── services/                  # Business logic layer
│   └── utils/                     # JSON response helpers
├── migrations/
│   ├── 000001_posts_blog_table.up.sql
│   └── 000001_posts_blog_table.down.sql
├── .env.example
├── go.mod
└── LICENSE
```

**Architecture flow:** `Handler -> Service -> Repository -> PostgreSQL`

## Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL

### Installation

```bash
git clone https://github.com/your-username/blog-api.git
cd blog-api
```

### Configuration

Create a `.env` file in the project root:

```
DATABASE_URL=postgres://user:password@localhost:5432/database_name?sslmode=disable
PORT=8080
```

| Variable | Required | Default | Description |
|---|---|---|---|
| `DATABASE_URL` | Yes | — | PostgreSQL connection string |
| `PORT` | No | `8080` | Server listen port |

### Database Setup

Run the migration SQL files manually against your PostgreSQL database:

```bash
psql -U postgres -d blog_manager -f migrations/000001_posts_blog_table.up.sql
```

This creates the `posts` table with the following schema:

```sql
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(150) NOT NULL,
    content VARCHAR(255),
    category VARCHAR(100),
    tags TEXT[]
);

CREATE INDEX idx_posts_category ON posts(category);
CREATE INDEX idx_posts_tags ON posts USING GIN(tags);
```

### Run

```bash
go run ./cmd/api/main.go
```

The server starts at `http://localhost:<PORT>`.

### Build

```bash
go build -o blog-api ./cmd/api/
```

## API Reference

### Base URL

```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/posts` | Create a new post |
| `GET` | `/posts` | List all posts / Search with query params |
| `GET` | `/posts/{id}` | Get a post by ID |
| `PATCH` | `/posts/{id}` | Partially update a post |
| `DELETE` | `/posts/{id}` | Delete a post |

---

### `POST /posts` — Create Post

**Request:**

```json
{
  "title": "My First Post",
  "content": "Some content here",
  "category": "tech",
  "tags": ["go", "api"]
}
```

Only `title` is required (max 150 chars). All other fields are optional.

**Response — `201 Created`:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "My First Post",
  "content": "Some content here",
  "category": "tech",
  "tags": ["go", "api"]
}
```

| Status | Reason |
|---|---|
| `400` | Invalid JSON or empty `title` |

---

### `GET /posts` — List All Posts

Returns an array of all posts. Returns `[]` if none exist.

**Response — `200 OK`:**

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "My First Post",
    "content": "Some content here",
    "category": "tech",
    "tags": ["go", "api"]
  }
]
```

#### Search / Filter

Pass query parameters to search and filter:

| Param | Description |
|---|---|
| `q` | Full-text search across `title` and `content` (case-insensitive) |
| `category` | Filter by category (prefix match) |
| `tag` | Filter by tag |

Example:

```
GET /posts?q=api&category=tech&tag=go
```

---

### `GET /posts/{id}` — Get Post by ID

**Response — `200 OK`:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "My First Post",
  "content": "Some content here",
  "category": "tech",
  "tags": ["go", "api"]
}
```

| Status | Reason |
|---|---|
| `500` | Post not found or invalid ID |

---

### `PATCH /posts/{id}` — Partial Update

All fields are optional — only the fields you send will be updated.

```json
{
  "title": "Updated Title",
  "tags": ["go", "api", "update"]
}
```

**Behavior:**

- `title`, `content`, `category`: sending an empty string (`""`) keeps the original value.
- `tags`: sending an empty array (`[]`) clears all tags.
- Unsent fields remain unchanged.

**Response — `200 OK`:**

Returns the input data (not the updated DB row). The `id` field comes back empty since it's not modified.

```json
{
  "id": "",
  "title": "Updated Title",
  "content": "",
  "category": "",
  "tags": ["go", "api", "update"]
}
```

| Status | Reason |
|---|---|
| `400` | Invalid JSON or all fields empty/missing |

---

### `DELETE /posts/{id}` — Delete Post

**Response — `200 OK`:**

Returns the deleted post's ID with all other fields empty.

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "",
  "content": "",
  "category": "",
  "tags": null
}
```

| Status | Reason |
|---|---|
| `500` | Post not found or invalid ID |

---

## Error Handling

All errors are returned as plain JSON strings:

```json
"error message here"
```

| Status | When |
|---|---|
| `200 OK` | Successful operation (list, get, update, delete) |
| `201 Created` | Post created successfully |
| `400 Bad Request` | Invalid JSON, missing required field, or invalid data |
| `500 Internal Server Error` | Server-side error (database, etc.) |

## License

[MIT](LICENSE)
