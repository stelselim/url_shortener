
# 📌 URL Shortener API Documentation

This project is a simple URL shortener service built with **Go (Echo)**. It allows users to generate short links, redirect to original URLs, view analytics, and manage shortened URLs.

<br>

## 🔗 API Endpoints

| Method   | Endpoint              | Purpose                             |
| -------- | --------------------- | ----------------------------------- |
| `POST`   | `/shorten`            | Create a new short URL              |
| `GET`    | `/:shortCode`         | Redirect to original URL            |
| `GET`    | `/stats/:shortCode`   | Get stats (hits, created\_at, etc.) |
| `DELETE` | `/shorten/:shortCode` | Delete a shortened URL              |


### 1. `POST /shorten`

Create a new short URL.

**Request Body:**
```json
{
  "original_url": "https://example.com/some/very/long/url"
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

### 2. `GET /:shortCode`

Redirect to the original long URL.

**Example:**
```
GET /abc123
```

**Response:**  
🔁 HTTP 302 Redirect to the original URL.

---

### 3. `GET /stats/:shortCode`

Get usage statistics of a short URL.

**Example:**
```
GET /stats/abc123
```

**Response:**
```json
{
  "original_url": "https://example.com/some/very/long/url",
  "short_code": "abc123",
  "created_at": "2025-05-14T10:00:00Z",
  "clicks": 42
}
```

---

### 4. `DELETE /shorten/:shortCode`

Delete a shortened URL.

**Example:**
```
DELETE /shorten/abc123
```

**Response:**
```json
{
  "message": "Short URL deleted successfully"
}
```

---

## 🚀 List of Required Concepts
- [X] Controller, Router, Services architecture
- [X] Echo library for building RESTful APIs.  
- [X] Create a Base Model Error and Response handling Errors, and HTTP Requests.
- [X] Firestore usage as a NoSQL DB.
- [X] Inject important values as environment variables.
- [X] No Deployment, Local Running
- [ ] Integration Tests
