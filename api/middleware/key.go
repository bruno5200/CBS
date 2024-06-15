package middleware

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
)

const (
	headerKey = "X-API-KEY"
	pagosKey  = "JDJhJDA0JGQ5VkpkeVNBeHJ2Nk82ai5QNzRHanUuT0VZUkhLMHlDZU81SkQvcmlmQUlwODRKRzdkQUJx"
	pagosHash = "$2a$04$d9VJdySAxrv6O6j.P74Gju.OEYRHK0yCeO5JD/rifAIp84JG7dABq"
)

func ApiKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// decode base64 string
		key, err := apiKey(utils.CopyString(c.Get(headerKey)))

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unauthorized"})
		}

		if key != pagosHash {
			return c.Status(fiber.StatusNotFound).SendFile("./www/unauthorized.html")
		}
		return c.Next()
	}
}

func apiKey(auth string) (hash string, err error) {
	key := auth[len("Key "):]
	// Split the base64 encoded string
	decoded, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
