Here are some improvements to your README for the `go-url-shortner` project:

---

# go-url-shortner

### Prerequisites

- Docker
- Git

### Steps to Run

```bash
git clone https://github.com/J0SAL/go-url-shortner.git
docker-compose up -d
```

### Example

#### Generating a Shortened URL

```bash
curl --location 'http://localhost:3000/api/v1' \
--header 'Content-Type: application/json' \
--data '{
    "url": "<URL>"
}'
```

#### Using a Shortened URL

```bash
http://localhost:3000/<url-id>
```

### Schema

#### Request Schema

```go
type request struct {
    url               string    //  URL to be shortened
    short             string    // Custom user ID (optional)
    expiry            integer   // Expiration time in hours
}
```

#### Response Schema

```go
type response struct {
    url               string    // BASE
    short             string    // shorten URL
    expiry            integer   // Expiration time in hours
    rate_limit        int       // Remaining requests that can be made
    rate_limit_reset  time      // Time left to reset the rate limit
}
```
