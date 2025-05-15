# URL Shortener

This project is a Go-based URL shortening service.  It has the ability to shorten lengthy URLs, store them in a database, and use a short code to get the original URLs.  The application runs in a containerized environment with Docker and Docker Compose and uses Redis and MySQL for caching and storage.

## Features

 * **URL Shortening**: For any lengthy URL, create a special short code.
 * **Redirection**: Send clients back to the original URL after using the short code.

 * **Persistent Storage**: Store URL mappings in MySQL.
 * **Caching**: For quicker lookups, use Redis.
 * **Static Web Interface**: A simple front-end to interact with the service.
 * **Portable and Scalable**: Docker and Docker Compose enable containerization.


## Technologies Used

* **Backend**: Go with Fiber framework
* **Database**: MySQL
* **Cache**: Redis
* **Frontend**: HTML, CSS, JavaScript
* **Containerization**: Docker

## Project Structure

```plaintext
URL-Shortener/
|-- Dockerfile               # Dockerfile for building the application image
|-- docker-compose.yml       # Docker Compose file for setting up the environment
|-- main.go                  # Entry point of the application
|-- db.go                    # Database initialization and table creation
|-- base62.go                # Base62 encoding logic
|-- public/                  # Static files for the web interface
|   |-- index.html           # Frontend HTML file
|   |-- style.css            # CSS styles for the frontend
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

1. Enter a long URL into the input box on the web interface.
2. Click the **Shorten** button to generate a short URL.
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

1. The user submits a long URL through the web interface or API endpoint.
2. A unique short code is generated using Base62 encoding.
3. The long URL and short code are stored in MySQL.
4. Redis is used for caching the mappings to improve lookup speed.
5. Users are redirected to the original URL by accessing the short code.
