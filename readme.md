# URL SHORTENER
The Project provide the URL shortening service.It shorten the lengthy URL, store them in database and use a short code to get the original URLs. It is a Golang based application that runs in a containerized environment with Docker. It uses Redis for caching and MySQL for the storage.

## Features
* **URL Shortening** : It  make long URLs in to short codes.
* **Redirection** : It sends the client to the original URL through the new generated short codes.
* **Persistent storage** : All the URLs and there short codes are mapped in the MySQl database.
* **Caching** : For quick lookups the cache data is stored in the Redis.
* **Protable and Scalable**: Docker is implemented to enable containerization.
* **Static Web Interface**: A Static web page front-end to access the services.

## Technologies Used
* **Backend**: Go with Fiber framework
* **Database**: MySQL
* **Cache**: Redis
* **Frontend**: HTML, CSS, JavaScript
* **Containerization**: Docker

## Project Structure

```plaintext
URL-Shortener/
|-- Dockerfile               # Dockerfile for to create the application image
|-- docker-compose.yml       # Docker Compose file for configiration the environment
|-- main.go                  # Entry point of the application
|-- db.go                    # Database initialization and table creation
|-- base62.go                # Base62 encoding logic
|-- public/                  # Static files of front-end
|   |-- index.html           # Frontend HTML file
|   |-- style.css            # CSS styles for the web interface
|   |-- script.js            # JavaScript functionality for the frontend
|-- go.mod                   # Go module dependencies
|-- go.sum                   # Checksums for Go modules
```

## Prerequisites

* Docker
* Docker Compose

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mhatrejeets/URL-Shortener.git
   cd URL-Shortener
   ```

2. Build and start the containers:

   ```bash
   docker-compose up --build
   ```

3. Access the application at:

   ```
   http://localhost:3000
   ```

## Frontend Usage

1. Enter a long URL into the input box on the front-end.
2. Click the **Shorten** button to generate a special short code mapped from the lengthy URLs.
3. The resulting short URL will be displayed below the input box.
4. Use the **Copy** button to copy the short URL to your clipboard.

## API Endpoints

* **POST /shorten**: Shorten a long URL.

  * Request body:

    ```json
    {
      "url": "https://example.com"
    }
    ```
  * Response:

    ```json
    {
      "short_url": "http://localhost:3000/abc123"
    }
    ```

* **GET /\:shortcode**: Redirect to the original URL associated with the shortcode.

## Environment Variables

The following environment variables can be configured in the `docker-compose.yml` file:

* `DB_HOST`: Hostname for the MySQL database (default: `mysql` for Docker).
* `REDIS_HOST`: Hostname for Redis (default: `redis` for Docker).
* `MYSQL_ROOT_PASSWORD`: Root password for MySQL.
* `MYSQL_DATABASE`: Database name for the application.
* `MYSQL_USER`: Username for MySQL.
* `MYSQL_PASSWORD`: Password for the specified user.
## Database Schema

The MySQL database table `urls` is created automatically with the following schema:

```sql
CREATE TABLE IF NOT EXISTS urls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    long_url VARCHAR(512) NOT NULL UNIQUE,
    short_code VARCHAR(7) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## How It Works

1. The user enter the lengthy URLs through the web interface or the API end-points..
2. Then a unique short code is generated based on Base62 encoding.
3. The long URL and short code are mapped and are stored in MySQL.
4. Redis is used for caching and to improve lookup speed.
5. Users are redirected to the original URL by accessing the short code.
