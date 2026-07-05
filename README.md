# Simple Blog

A simple blog REST API written in Go (Echo + GORM + MySQL) supporting user registration/login, blog posts, and comments.

## Running with Docker Compose

### Prerequisites

- Docker
- Docker Compose

### Start the app

```bash
docker compose build
docker compose up
```

This brings up three services:

| Service   | Description                                                                 |
|-----------|-------------------------------------------------------------------------------|
| `mysql`   | MySQL 8.0, exposed on host port `13306` (configurable via `DB_PORT` in `.env`) |
| `migrate` | Runs all pending SQL migrations in `database/migrations`, then exits          |
| `app`     | The API server, exposed on `http://localhost:8080`                            |

The `app` service waits for `mysql` to report healthy and for `migrate` to complete successfully before starting.

To run it in the background:

```bash
docker compose up -d
```

To stop and remove the containers:

```bash
docker compose down
```

To wipe the database volume as well:

```bash
docker compose down -v
```

### Configuration

Environment variables are read from `.env` (see `.env.example`). Defaults used by `docker-compose.yml` if unset:

| Variable          | Default                       | Description                          |
|-------------------|--------------------------------|---------------------------------------|
| `DB_HOST`         | `mysql` (fixed, inside compose) | MySQL host                            |
| `DB_PORT`         | `13306`                        | Host-mapped MySQL port                |
| `DB_DATABASE`     | `simple_blog`                  | MySQL database name                   |
| `DB_USERNAME`     | `admin`                        | MySQL user                            |
| `DB_PASSWORD`     | `admin123`                     | MySQL password                        |
| `JWT_SECRET`      | `change-me-in-development`     | Secret used to sign JWTs              |
| `JWT_TTL_MINUTES` | `60`                           | Access token lifetime, in minutes     |

> Note: port `13306` is used instead of the default `3306` to avoid clashing with a locally installed MySQL. If that port is already taken on your machine, change `DB_PORT` in `.env`.

## Authentication

Most endpoints require a JWT access token, obtained via `POST /login`, sent as:

```
Authorization: Bearer <access_token>
```

The following endpoints are public (no token required):

- `POST /register`
- `POST /login`
- `GET /posts`
- `GET /posts/:id`
- `GET /posts/:id/comments`
- `POST /posts/:id/comments` (token optional — if provided, the comment is attributed to that user)

All other endpoints return `401 Unauthorized` without a valid token.

## API Endpoints

All responses share the envelope:

```json
{
  "message": "string",
  "data": "...",
  "metadata": "...",
  "error": "..."
}
```

### Users

#### Register

`POST /register`

Request body:

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "Sup3r$ecret"
}
```

- `name`, `email`, `password` required; `password` must be 8-72 characters and combine at least 3 of: uppercase, lowercase, digit, special character (and must not be a common/personal password).
- Returns `201 Created` on success, `409 Conflict` if the email already exists, `422` for weak passwords.

#### Login

`POST /login`

Request body:

```json
{
  "email": "jane@example.com",
  "password": "Sup3r$ecret"
}
```

Response `200 OK`:

```json
{
  "message": "Successfully logged in",
  "data": {
    "access_token": "<jwt>",
    "expires_at": "2026-07-05T12:00:00Z"
  }
}
```

Returns `401 Unauthorized` for invalid credentials.

### Posts

#### Create post

`POST /posts` — **requires auth**

Request body:

```json
{
  "title": "My first post",
  "content": "Hello world"
}
```

Returns `201 Created` with the new post's `id`.

#### List posts

`GET /posts` — public

Query params (all optional):

| Param      | Default | Description                       |
|------------|---------|-----------------------------------|
| `page`     | `1`     | Page number                       |
| `per_page` | `5`     | Items per page                    |
| `order`    | -       | Sort direction (e.g. `asc`/`desc`)|
| `order_by` | -       | Column to sort by                 |

Response `200 OK`:

```json
{
  "message": "Successfully fetched blog post list",
  "data": [
    {
      "id": 1,
      "title": "My first post",
      "author_id": 1,
      "author_name": "Jane Doe",
      "created_at": "2026-07-05T12:00:00Z",
      "updated_at": "2026-07-05T12:00:00Z"
    }
  ],
  "metadata": {
    "page": 1,
    "per_page": 5,
    "total_page": 1,
    "total_data": 1
  }
}
```

#### Get post detail

`GET /posts/:id` — public

Response `200 OK`:

```json
{
  "message": "Successfully fetched blog post detail",
  "data": {
    "id": 1,
    "title": "My first post",
    "content": "Hello world",
    "author_name": "Jane Doe",
    "created_at": "2026-07-05T12:00:00Z",
    "updated_at": "2026-07-05T12:00:00Z"
  }
}
```

Returns `400 Bad Request` if the post doesn't exist.

#### Update post

`PUT /posts/:id` — **requires auth (author only)**

Request body:

```json
{
  "title": "Updated title",
  "content": "Updated content"
}
```

Returns `200 OK` on success, `403 Forbidden` if the requester isn't the post's author, `400 Bad Request` if the post doesn't exist.

#### Delete post

`DELETE /posts/:id` — **requires auth (author only)**

Returns `200 OK` on success, `403 Forbidden` if the requester isn't the post's author, `400 Bad Request` if the post doesn't exist.

### Comments

#### Add comment

`POST /posts/:id/comments` — public (token optional)

Request body:

```json
{
  "content": "Great post!"
}
```

Returns `201 Created` with the new comment's `id`, `400 Bad Request` if the post doesn't exist.

#### List comments

`GET /posts/:id/comments` — public

Query params (all optional): same `page`, `per_page`, `order`, `order_by` as the post list endpoint.

Response `200 OK`:

```json
{
  "message": "Successfully fetched post's comment list",
  "data": [
    {
      "id": 1,
      "post_id": 1,
      "author_id": 1,
      "author_name": "Jane Doe",
      "content": "Great post!",
      "created_at": "2026-07-05T12:00:00Z"
    }
  ],
  "metadata": {
    "page": 1,
    "per_page": 5,
    "total_page": 1,
    "total_data": 1
  }
}
```

## Local development (without Docker)

Prerequisites: Go 1.25+, MySQL 8.0, [Air](https://github.com/air-verse/air) for hot reload.

```bash
go install github.com/air-verse/air@latest
air
```

Make sure `.env` points at a running MySQL instance and migrations in `database/migrations` have been applied.
