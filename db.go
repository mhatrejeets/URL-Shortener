package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	DB  *sql.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func initDB() {
	dsn := "username:password@tcp(mysql:3306)/shortnerdb"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INT AUTO_INCREMENT PRIMARY KEY,
		long_url VARCHAR(512) NOT NULL UNIQUE,
		short_code VARCHAR(7) UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal("Table creation error:", err)
	}
}

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		//use redis: for docker
		Addr: "redis:6379",

	})
	if err := RDB.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
