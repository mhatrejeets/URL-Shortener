package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func createShortURL(c *fiber.Ctx) error {
	type Request struct {
		URL string `json:"url"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid URL or input")
	}

	
	var existingCode string
	err := DB.QueryRow("SELECT short_code FROM urls WHERE long_url = ?", body.URL).Scan(&existingCode)
	if err == nil {
		
		return c.JSON(fiber.Map{"short_url": c.BaseURL() + "/" + existingCode})
	}

	
	res, err := DB.Exec("INSERT INTO urls (long_url) VALUES (?)", body.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("DB Error (Not Inserted long url.)")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failure in fetching the ID")
	}

	code := encodeBase62(id)

	_, err = DB.Exec("UPDATE urls SET short_code = ? WHERE id = ?", code, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("DB Error (short_code not updated)")
	}

	err = RDB.Set(Ctx, code, body.URL, 1*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Redis error")
	}

	return c.JSON(fiber.Map{"short_url": c.BaseURL() + "/" + code})
}

func RedirectToLongURL(c *fiber.Ctx) error {
	code := c.Params("shortcode")
	longURL, err := RDB.Get(Ctx, code).Result()
	if err == redis.Nil {
	row := DB.QueryRow("SELECT long_url FROM urls WHERE short_code = ?", code)
	err = row.Scan(&longURL)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("URL Not Found")
	}
	RDB.Set(Ctx, code, longURL, 1*time.Hour)
	}
	

	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}

	return c.Redirect(longURL, fiber.StatusFound)

}
