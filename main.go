package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}

	
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Log to file
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Warn("Failed to log to file, using default stderr")
	}

	initDB()
	initRedis()

	app := fiber.New()

	app.Static("/", "./public")
	app.Post("/shorten", createShortURL)
	app.Get("/:shortcode", RedirectToLongURL)

	logrus.Info("Server is starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
	defer DB.Close()
}
