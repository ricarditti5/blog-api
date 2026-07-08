# Blog API — Pré-documentação

## Stack

- **Linguagem:** Go 1.26.4
- **Framework:** Nenhum (`net/http` nativo com Go 1.22+ pattern routing)
- **Base de dados:** PostgreSQL via `pgx/v5` (pgxpool)
- **Config:** `viper` + `gotenv`

---

## Base URL

```
http://localhost:8080
```

A porta é definida pela variável de ambiente `PORT` (default: `8080`).

---

## Endpoints

| Método | Rota | Descrição |
|---|---|---|
| `POST` | `/posts` | Criar um post |
| `GET` | `/posts` | Listar todos os posts |
| `GET` | `/posts/{id}` | Obter um post por ID |
| `PATCH` | `/posts/{id}` | Atualizar parcialmente um post |
| `DELETE` | `/posts/{id}` | Eliminar um post |

---

## Modelo — Post

```json
{
  "id": "uuid",
  "title": "string",
  "content": "string",
  "category": "string",
  "tags": ["string"]
}
```

| Campo | Tipo | Obrigatório | Descrição |
|---|---|---|---|
| `id` | `string` (UUID) | gerado pela BD | Chave primária, auto-gerado |
| `title` | `string` | sim | Título do post (max 150) |
| `content` | `string` | não | Conteúdo (max 255) |
| `category` | `string` | não | Categoria (max 100) |
| `tags` | `array<string>` | não | Tags do post (TEXT[]) |

---

## Endpoints — Detalhe

### `POST /posts` — Criar post

**Request body:**

```json
{
  "title": "Meu Post",
  "content": "Conteúdo aqui",
  "category": "tech",
  "tags": ["go", "api"]
}
```

Apenas `title` é obrigatório. Os restantes campos são opcionais.

**Response — `201 Created`:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Meu Post",
  "content": "Conteúdo aqui",
  "category": "tech",
  "tags": ["go", "api"]
}
```

**Errors:**

| Status | Motivo |
|---|---|
| `400` | JSON inválido ou `title` vazio |

---

### `GET /posts` — Listar todos os posts

**Response — `200 OK`:**

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Meu Post",
    "content": "Conteúdo aqui",
    "category": "tech",
    "tags": ["go", "api"]
  }
]
```

Devolve um array vazio `[]` se não existirem posts.

---

### `GET /posts/{id}` — Obter post por ID

**Response — `200 OK`:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Meu Post",
  "content": "Conteúdo aqui",
  "category": "tech",
  "tags": ["go", "api"]
}
```

**Errors:**

| Status | Motivo |
|---|---|
| `500` | Post não encontrado ou ID inválido |

---

### `PATCH /posts/{id}` — Atualizar parcialmente um post

**Request body** (todos os campos opcionais — só os enviados são alterados):

```json
{
  "title": "Título Atualizado",
  "tags": ["go", "api", "update"]
}
```

**Comportamento:** Campos não enviados no JSON mantêm o valor original na base de dados.

- `title`, `content`, `category`: se enviados como string vazia (`""`), o valor original mantém-se (não é possível limpar para vazio via PATCH).
- `tags`: se enviado como array vazia (`[]`), as tags são limpas.

**Response — `200 OK`:**

Devolve os dados enviados (não os atuais da BD).

```json
{
  "id": "",
  "title": "Título Atualizado",
  "content": "",
  "category": "",
  "tags": ["go", "api", "update"]
}
```

> Nota: o campo `id` no response vem vazio porque não é alterado pela query.

**Errors:**

| Status | Motivo |
|---|---|
| `400` | JSON inválido ou todos os campos vazios/não enviados |

---

### `DELETE /posts/{id}` — Eliminar um post

**Response — `200 OK`:**

Devolve o objeto com o ID do post eliminado (restantes campos vazios).

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "",
  "content": "",
  "category": "",
  "tags": null
}
```

**Errors:**

| Status | Motivo |
|---|---|
| `500` | Post não encontrado ou ID inválido |

---

## Error Handling

Todos os erros seguem o formato:

```json
"mensagem de erro"
```

(Enviado como string JSON, não como objeto `{"error": "..."}`.)

### Status codes utilizados

| Status | Quando usado |
|---|---|
| `200 OK` | Operação bem-sucedida (list, get, update, delete) |
| `201 Created` | Post criado com sucesso |
| `400 Bad Request` | JSON inválido, campo obrigatório vazio, ou dados inválidos |
| `500 Internal Server Error` | Erro do servidor (BD, etc.) |

---

## Arquitetura do Código

```
cmd/api/main.go          → Entry point
internal/config/         → Config (.env + viper)
internal/handlers/       → HTTP handlers
internal/services/       → Business logic
internal/repository/     → Database queries
internal/models/         → Data structures
internal/routes/         → Route registration
internal/utils/          → Response helpers
```

Fluxo: `Handler → Service → Repository → PostgreSQL`

---

## Base de Dados

### Tabela `posts`

```sql
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(150) NOT NULL,
    content varchar(255),
    category varchar(100),
    tags TEXT[]
);
```

### Índices

```sql
CREATE INDEX idx_posts_category ON posts(category);
CREATE INDEX idx_posts_tags ON posts USING GIN(tags);
```

---

## Variáveis de Ambiente (`.env`)

```
DATABASE_URL=postgres://user:password@host:port/dbname?sslmode=disable
PORT=8080
```
