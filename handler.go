package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func createShortURL(c *fiber.Ctx) error {
	type Request struct {
		URL string `json:"url"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		logrus.WithError(err).Warn("Invalid input to shorten")
		return c.Status(fiber.StatusBadRequest).SendString("Invalid URL or input")
	}

	var existingCode string
	err := DB.QueryRow("SELECT short_code FROM urls WHERE long_url = ?", body.URL).Scan(&existingCode)
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"url":        body.URL,
			"short_code": existingCode,
		}).Info("URL already exists in DB")
		return c.JSON(fiber.Map{"short_url": c.BaseURL() + "/" + existingCode})
	}

	res, err := DB.Exec("INSERT INTO urls (long_url) VALUES (?)", body.URL)
	if err != nil {
		logrus.WithError(err).Error("Failed to insert long URL")
		return c.Status(fiber.StatusInternalServerError).SendString("DB Error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.WithError(err).Error("Failed to get last insert ID")
		return c.Status(fiber.StatusInternalServerError).SendString("Failure in fetching the ID")
	}

	code := encodeBase62(id)

	_, err = DB.Exec("UPDATE urls SET short_code = ? WHERE id = ?", code, id)
	if err != nil {
		logrus.WithError(err).Error("Failed to update short code")
		return c.Status(fiber.StatusInternalServerError).SendString("DB Error (short_code not updated)")
	}

	err = RDB.Set(Ctx, code, body.URL, 1*time.Hour).Err()
	if err != nil {
		logrus.WithError(err).Error("Redis SET failed")
		return c.Status(fiber.StatusInternalServerError).SendString("Redis error")
	}

	logrus.WithFields(logrus.Fields{
		"short_code": code,
		"url":        body.URL,
	}).Info("Short URL created")

	return c.JSON(fiber.Map{"short_url": c.BaseURL() + "/" + code})
}

func RedirectToLongURL(c *fiber.Ctx) error {
	code := c.Params("shortcode")

	longURL, err := RDB.Get(Ctx, code).Result()
	if err == redis.Nil {
		row := DB.QueryRow("SELECT long_url FROM urls WHERE short_code = ?", code)
		err = row.Scan(&longURL)
		if err != nil {
			logrus.WithField("shortcode", code).Warn("Shortcode not found in DB or Redis")
			return c.Status(fiber.StatusNotFound).SendString("URL Not Found")
		}
		RDB.Set(Ctx, code, longURL, 1*time.Hour)
		logrus.WithField("shortcode", code).Info("Cache miss, loaded from DB")
	} else if err != nil {
		logrus.WithError(err).Error("Redis GET error")
		return c.Status(fiber.StatusInternalServerError).SendString("Redis error")
	}

	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}

	logrus.WithFields(logrus.Fields{
		"shortcode": code,
		"long_url":  longURL,
	}).Info("Redirecting to long URL")

	return c.Redirect(longURL, fiber.StatusFound)
}
