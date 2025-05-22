package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	DB  *sql.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func initDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "username:password@tcp(mysql:3306)/shortnerdb"
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open DB connection")
	}

	if err = DB.Ping(); err != nil {
		logrus.WithError(err).Fatal("Failed to ping DB")
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INT AUTO_INCREMENT PRIMARY KEY,
		long_url VARCHAR(512) NOT NULL UNIQUE,
		short_code VARCHAR(7) UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create 'urls' table")
	}

	logrus.Info("Database initialized successfully")
}

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		logrus.WithError(err).Fatal("Failed to connect to Redis")
	}

	logrus.Info("Redis initialized successfully")
}
