
# URL Shortener

URL Shortener is a simple URL shortening service written in Go. It provides a way to convert long URLs into short, manageable links. This project is intended for educational purposes and demonstrates the use of Go for building web services.

## Features

- Shorten long URLs to short aliases
- Redirect short URLs to the original long URLs
- Basic in-memory storage for URL mappings

## Getting Started

### Prerequisites

- Go 1.16 or later
- Git

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/surl-shortener.git
   cd surl-shortener
   ```

2. Build the project:
   ```sh
   go build
   ```

3. Run the project:
   ```sh
   ./url_shortener
   ```

### Usage

1. Start the server:
   ```sh
   go run main.go
   ```

2. Shorten a URL:
   Send a POST request to `/shorten` with a JSON body containing the long URL.
   ```sh
   curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten
   ```

3. Redirect to the long URL:
   Access the short URL in the browser or via curl.
   ```sh
   curl http://localhost:8080/{shortURL}
   ```


### Example

1. Start the server:
   ```sh
   go run main.go
   ```

2. Shorten a URL:
   ```sh
   curl -X POST -H "Content-Type: application/json" -d '{"url":"https://golang.org"}' http://localhost:8080/shorten
   ```

3. The response will be a shortened URL:
   ```json
   {
       "short_url": "http://localhost:8080/abc123"
   }
   ```

4. Redirect to the original URL:
   ```sh
   curl http://localhost:8080/abc123
   ```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request with any improvements or bug fixes.

###